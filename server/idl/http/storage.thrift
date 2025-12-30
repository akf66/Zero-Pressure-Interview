namespace go storage

include "../base/common.thrift"
include "../base/storage.thrift"

// ==================== 获取上传URL ====================
struct GetUploadUrlRequest {
    1: string file_name (api.query="file_name", api.vd="len($) > 0"),
    2: string file_type (api.query="file_type", api.vd="len($) > 0"),
    3: string content_type (api.query="content_type"),
}

struct GetUploadUrlResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: string upload_url (api.body="upload_url"),
    4: string file_key (api.body="file_key"),
    5: i64 expires_at (api.body="expires_at"),
}

// ==================== 确认上传 ====================
struct ConfirmUploadRequest {
    1: string file_key (api.body="file_key", api.vd="len($) > 0"),
    2: string file_type (api.body="file_type", api.vd="len($) > 0"),
}

struct ConfirmUploadResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: string file_url (api.body="file_url"),
}

// ==================== 获取下载URL ====================
struct GetDownloadUrlResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: string download_url (api.body="download_url"),
    4: i64 expires_at (api.body="expires_at"),
}

// ==================== 批量删除文件 ====================
struct BatchDeleteFilesRequest {
    1: list<string> file_keys (api.body="file_keys", api.vd="len($) > 0"),
}

struct BatchDeleteFilesResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: i32 deleted_count (api.body="deleted_count"),
}

// ==================== 获取文件信息 ====================
struct GetFileInfoResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: storage.FileInfo file_info (api.body="file_info"),
}

// ==================== 获取用户文件列表 ====================
struct GetUserFilesRequest {
    1: string file_type (api.query="file_type"),
    2: i32 page (api.query="page", api.vd="$ > 0"),
    3: i32 page_size (api.query="page_size", api.vd="$ > 0 && $ <= 100"),
}

struct GetUserFilesResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: list<storage.FileInfo> files (api.body="files"),
    4: common.PageResponse page (api.body="page"),
}

// ==================== 文件服务接口 ====================
service StorageService {
    GetUploadUrlResponse GetUploadUrl(1: GetUploadUrlRequest req) (api.get="/api/v1/file/upload-url"),
    ConfirmUploadResponse ConfirmUpload(1: ConfirmUploadRequest req) (api.post="/api/v1/file/confirm"),
    GetDownloadUrlResponse GetDownloadUrl(1: string file_key) (api.get="/api/v1/file/download-url", api.query="file_key"),
    GetFileInfoResponse GetFileInfo(1: string file_key) (api.get="/api/v1/file/info/:file_key", api.path="file_key"),
    GetUserFilesResponse GetUserFiles(1: GetUserFilesRequest req) (api.get="/api/v1/file/list"),
    common.NilResponse DeleteFile(1: string file_key) (api.delete="/api/v1/file/:file_key", api.path="file_key"),
    BatchDeleteFilesResponse BatchDeleteFiles(1: BatchDeleteFilesRequest req) (api.post="/api/v1/file/batch-delete"),
}