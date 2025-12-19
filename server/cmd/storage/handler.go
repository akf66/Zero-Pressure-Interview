package main

import (
	"context"
	storage "zpi/server/shared/kitex_gen/storage"
)

// StorageServiceImpl implements the last service interface defined in the IDL.
type StorageServiceImpl struct{}

// GetUploadUrl implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetUploadUrl(ctx context.Context, req *storage.GetUploadUrlReq) (resp *storage.GetUploadUrlResp, err error) {
	// TODO: Your code here...
	return
}

// ConfirmUpload implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) ConfirmUpload(ctx context.Context, req *storage.ConfirmUploadReq) (resp *storage.ConfirmUploadResp, err error) {
	// TODO: Your code here...
	return
}

// GetDownloadUrl implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetDownloadUrl(ctx context.Context, req *storage.GetDownloadUrlReq) (resp *storage.GetDownloadUrlResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteFile implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) DeleteFile(ctx context.Context, req *storage.DeleteFileReq) (resp *storage.DeleteFileResp, err error) {
	// TODO: Your code here...
	return
}

// GetFileInfo implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetFileInfo(ctx context.Context, req *storage.GetFileInfoReq) (resp *storage.GetFileInfoResp, err error) {
	// TODO: Your code here...
	return
}

// BatchDeleteFiles implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) BatchDeleteFiles(ctx context.Context, req *storage.BatchDeleteFilesReq) (resp *storage.BatchDeleteFilesResp, err error) {
	// TODO: Your code here...
	return
}

// GetUserFiles implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetUserFiles(ctx context.Context, req *storage.GetUserFilesReq) (resp *storage.GetUserFilesResp, err error) {
	// TODO: Your code here...
	return
}
