package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"
	"zpi/server/cmd/storage/config"
	"zpi/server/cmd/storage/initialize"
	"zpi/server/shared/dal/sqlentity"
	"zpi/server/shared/dal/sqlfunc"
	"zpi/server/shared/errno"
	base "zpi/server/shared/kitex_gen/base"
	storage "zpi/server/shared/kitex_gen/storage"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

// StorageServiceImpl implements the last service interface defined in the IDL.
type StorageServiceImpl struct {
	*StorageManager
}

type StorageManager struct {
	Query *sqlfunc.Query
}

// fileTypeToString 将文件类型枚举转换为字符串
func fileTypeToString(ft base.FileType) string {
	switch ft {
	case base.FileType_RESUME:
		return "resume"
	case base.FileType_AVATAR:
		return "avatar"
	case base.FileType_RECORDING:
		return "recording"
	default:
		return "other"
	}
}

// generateFileKey 生成文件存储key
func generateFileKey(userID int64, fileType base.FileType, fileName string) string {
	ext := filepath.Ext(fileName)
	uid := uuid.New().String()
	typeStr := fileTypeToString(fileType)
	return fmt.Sprintf("%s/%d/%s%s", typeStr, userID, uid, ext)
}

// GetUploadUrl implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetUploadUrl(ctx context.Context, req *storage.GetUploadUrlRequest) (resp *storage.GetUploadUrlResponse, err error) {
	resp = new(storage.GetUploadUrlResponse)

	// 参数校验
	if req.UserId <= 0 || req.FileName == "" || req.ContentType == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr)
		return resp, nil
	}

	// 生成文件key
	fileKey := generateFileKey(req.UserId, req.FileType, req.FileName)
	bucket := config.GlobalServerConfig.MinIOInfo.Bucket

	// 生成预签名上传URL，有效期15分钟
	expiry := 15 * time.Minute
	presignedURL, err := initialize.MinIOClient.PresignedPutObject(ctx, bucket, fileKey, expiry)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.UploadUrl = presignedURL.String()
	resp.FileKey = fileKey
	resp.ExpiresAt = time.Now().Add(expiry).Unix()
	return resp, nil
}

// ConfirmUpload implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) ConfirmUpload(ctx context.Context, req *storage.ConfirmUploadRequest) (resp *storage.ConfirmUploadResponse, err error) {
	resp = new(storage.ConfirmUploadResponse)

	// 参数校验
	if req.UserId <= 0 || req.FileKey == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr)
		return resp, nil
	}

	bucket := config.GlobalServerConfig.MinIOInfo.Bucket

	// 检查文件是否存在
	objInfo, err := initialize.MinIOClient.StatObject(ctx, bucket, req.FileKey, minio.StatObjectOptions{})
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.FileNotFound)
		return resp, nil
	}

	// 提取文件名
	parts := strings.Split(req.FileKey, "/")
	fileName := parts[len(parts)-1]
	ext := filepath.Ext(fileName)

	// 保存文件记录到数据库
	st := s.Query.Storage
	storageRecord := &sqlentity.Storage{
		UserID:      uint64(req.UserId),
		FileName:    fileName,
		FilePath:    req.FileKey,
		FileSize:    objInfo.Size,
		FileType:    objInfo.ContentType,
		FileExt:     &ext,
		StorageType: "minio",
		Bucket:      &bucket,
		ObjectKey:   &req.FileKey,
		Status:      1,
	}

	if err := st.WithContext(ctx).Create(storageRecord); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 生成访问URL
	endpoint := config.GlobalServerConfig.MinIOInfo.Endpoint
	useSSL := config.GlobalServerConfig.MinIOInfo.UseSSL
	protocol := "http"
	if useSSL {
		protocol = "https"
	}
	fileURL := fmt.Sprintf("%s://%s/%s/%s", protocol, endpoint, bucket, req.FileKey)

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.FileUrl = fileURL
	return resp, nil
}

// GetDownloadUrl implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetDownloadUrl(ctx context.Context, req *storage.GetDownloadUrlRequest) (resp *storage.GetDownloadUrlResponse, err error) {
	resp = new(storage.GetDownloadUrlResponse)

	// 参数校验
	if req.UserId <= 0 || req.FileKey == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr)
		return resp, nil
	}

	bucket := config.GlobalServerConfig.MinIOInfo.Bucket

	// 检查文件是否存在且属于该用户
	st := s.Query.Storage
	_, err = st.WithContext(ctx).Where(st.FilePath.Eq(req.FileKey), st.UserID.Eq(uint64(req.UserId)), st.Status.Eq(1)).First()
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.FileNotFound)
		return resp, nil
	}

	// 生成预签名下载URL，有效期1小时
	expiry := 1 * time.Hour
	presignedURL, err := initialize.MinIOClient.PresignedGetObject(ctx, bucket, req.FileKey, expiry, nil)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.DownloadUrl = presignedURL.String()
	resp.ExpiresAt = time.Now().Add(expiry).Unix()
	return resp, nil
}

// DeleteFile implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) DeleteFile(ctx context.Context, req *storage.DeleteFileRequest) (resp *storage.DeleteFileResponse, err error) {
	resp = new(storage.DeleteFileResponse)

	// 参数校验
	if req.UserId <= 0 || req.FileKey == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr)
		return resp, nil
	}

	// 检查文件是否存在且属于该用户
	st := s.Query.Storage
	storageRecord, err := st.WithContext(ctx).Where(st.FilePath.Eq(req.FileKey), st.UserID.Eq(uint64(req.UserId)), st.Status.Eq(1)).First()
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.FileNotFound)
		return resp, nil
	}

	bucket := config.GlobalServerConfig.MinIOInfo.Bucket

	// 从MinIO删除文件
	if err := initialize.MinIOClient.RemoveObject(ctx, bucket, req.FileKey, minio.RemoveObjectOptions{}); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 软删除数据库记录
	if _, err := st.WithContext(ctx).Where(st.ID.Eq(storageRecord.ID)).Delete(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFileInfo implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetFileInfo(ctx context.Context, req *storage.GetFileInfoRequest) (resp *storage.GetFileInfoResponse, err error) {
	resp = new(storage.GetFileInfoResponse)

	// 参数校验
	if req.FileKey == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr)
		return resp, nil
	}

	// 查询文件记录
	st := s.Query.Storage
	storageRecord, err := st.WithContext(ctx).Where(st.FilePath.Eq(req.FileKey), st.Status.Eq(1)).First()
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.FileNotFound)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.FileInfo = &base.FileInfo{
		FileKey:     storageRecord.FilePath,
		FileName:    storageRecord.FileName,
		ContentType: storageRecord.FileType,
		FileSize:    storageRecord.FileSize,
		CreatedAt:   storageRecord.CreatedAt.Unix(),
	}
	return resp, nil
}

// BatchDeleteFiles implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) BatchDeleteFiles(ctx context.Context, req *storage.BatchDeleteFilesRequest) (resp *storage.BatchDeleteFilesResponse, err error) {
	resp = new(storage.BatchDeleteFilesResponse)

	// 参数校验
	if req.UserId <= 0 || len(req.FileKeys) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr)
		return resp, nil
	}

	bucket := config.GlobalServerConfig.MinIOInfo.Bucket
	deletedCount := int32(0)

	st := s.Query.Storage

	for _, fileKey := range req.FileKeys {
		// 检查文件是否存在且属于该用户
		storageRecord, err := st.WithContext(ctx).Where(st.FilePath.Eq(fileKey), st.UserID.Eq(uint64(req.UserId)), st.Status.Eq(1)).First()
		if err != nil {
			continue
		}

		// 从MinIO删除文件
		if err := initialize.MinIOClient.RemoveObject(ctx, bucket, fileKey, minio.RemoveObjectOptions{}); err != nil {
			continue
		}

		// 软删除数据库记录
		if _, err := st.WithContext(ctx).Where(st.ID.Eq(storageRecord.ID)).Delete(); err != nil {
			continue
		}

		deletedCount++
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.DeletedCount = deletedCount
	return resp, nil
}

// GetUserFiles implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetUserFiles(ctx context.Context, req *storage.GetUserFilesRequest) (resp *storage.GetUserFilesResponse, err error) {
	resp = new(storage.GetUserFilesResponse)

	// 参数校验
	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr)
		return resp, nil
	}

	st := s.Query.Storage

	// 构建查询
	query := st.WithContext(ctx).Where(st.UserID.Eq(uint64(req.UserId)), st.Status.Eq(1))

	// 文件类型筛选
	if req.FileType != nil {
		typeStr := fileTypeToString(*req.FileType)
		query = query.Where(st.FilePath.Like(typeStr + "/%"))
	}

	// 统计总数
	total, err := query.Count()
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 分页查询
	page := int32(1)
	pageSize := int32(10)
	if req.Page != nil {
		if req.Page.Page > 0 {
			page = req.Page.Page
		}
		if req.Page.PageSize > 0 {
			pageSize = req.Page.PageSize
		}
	}
	offset := int((page - 1) * pageSize)

	records, err := query.Order(st.CreatedAt.Desc()).Offset(offset).Limit(int(pageSize)).Find()
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 转换为响应格式
	files := make([]*base.FileInfo, 0, len(records))
	for _, r := range records {
		files = append(files, &base.FileInfo{
			FileKey:     r.FilePath,
			FileName:    r.FileName,
			ContentType: r.FileType,
			FileSize:    r.FileSize,
			CreatedAt:   r.CreatedAt.Unix(),
		})
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Files = files
	resp.Page = &base.PageResponse{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
	return resp, nil
}

// HealthCheck implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) HealthCheck(ctx context.Context) (resp *base.HealthCheckResponse, err error) {
	resp = &base.HealthCheckResponse{
		Status: "ok",
	}
	return resp, nil
}
