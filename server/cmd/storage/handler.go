package main

import (
	"context"
	storage "zpi/server/shared/kitex_gen/storage"
)

// StorageServiceImpl implements the last service interface defined in the IDL.
type StorageServiceImpl struct{}

// GetUploadUrl implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetUploadUrl(ctx context.Context, req *storage.GetUploadUrlRequest) (resp *storage.GetUploadUrlResponse, err error) {
	// TODO: Your code here...
	return
}

// ConfirmUpload implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) ConfirmUpload(ctx context.Context, req *storage.ConfirmUploadRequest) (resp *storage.ConfirmUploadResponse, err error) {
	// TODO: Your code here...
	return
}

// GetDownloadUrl implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetDownloadUrl(ctx context.Context, req *storage.GetDownloadUrlRequest) (resp *storage.GetDownloadUrlResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteFile implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) DeleteFile(ctx context.Context, req *storage.DeleteFileRequest) (resp *storage.DeleteFileResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFileInfo implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetFileInfo(ctx context.Context, req *storage.GetFileInfoRequest) (resp *storage.GetFileInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// BatchDeleteFiles implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) BatchDeleteFiles(ctx context.Context, req *storage.BatchDeleteFilesRequest) (resp *storage.BatchDeleteFilesResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserFiles implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) GetUserFiles(ctx context.Context, req *storage.GetUserFilesRequest) (resp *storage.GetUserFilesResponse, err error) {
	// TODO: Your code here...
	return
}
