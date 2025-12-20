namespace go storage

include "../base/common.thrift"
include "../base/storage.thrift"

// ==================== 获取上传URL ====================
struct GetUploadUrlRequest {
    1: i64 user_id
    2: string file_name                // 文件名
    3: storage.FileType file_type      // 文件类型：resume/avatar/recording
    4: string content_type             // MIME类型，如：application/pdf
}

struct GetUploadUrlResponse {
    1: common.BaseResponse base_resp
    2: string upload_url     // 预签名上传URL
    3: string file_key       // 文件存储key
    4: i64 expires_at        // 过期时间戳（秒）
}

// ==================== 确认上传 ====================
struct ConfirmUploadRequest {
    1: i64 user_id
    2: string file_key                 // 文件存储key
    3: storage.FileType file_type      // 文件类型
}

struct ConfirmUploadResponse {
    1: common.BaseResponse base_resp
    2: string file_url       // 文件访问URL
}

// ==================== 获取下载URL ====================
struct GetDownloadUrlRequest {
    1: i64 user_id
    2: string file_key       // 文件存储key
}

struct GetDownloadUrlResponse {
    1: common.BaseResponse base_resp
    2: string download_url   // 预签名下载URL
    3: i64 expires_at        // 过期时间戳（秒）
}

// ==================== 删除文件 ====================
struct DeleteFileRequest {
    1: i64 user_id
    2: string file_key       // 文件存储key
}

struct DeleteFileResponse {
    1: common.BaseResponse base_resp
}

// ==================== 获取文件信息 ====================
struct GetFileInfoRequest {
    1: string file_key       // 文件存储key
}

struct GetFileInfoResponse {
    1: common.BaseResponse base_resp
    2: storage.FileInfo file_info
}

// ==================== 批量删除文件 ====================
struct BatchDeleteFilesRequest {
    1: i64 user_id
    2: list<string> file_keys
}

struct BatchDeleteFilesResponse {
    1: common.BaseResponse base_resp
    2: i32 deleted_count     // 成功删除的数量
}

// ==================== 获取用户文件列表 ====================
struct GetUserFilesRequest {
    1: i64 user_id
    2: optional storage.FileType file_type  // 文件类型筛选
    3: common.PageRequest page
}

struct GetUserFilesResponse {
    1: common.BaseResponse base_resp
    2: list<storage.FileInfo> files
    3: common.PageResponse page
}

// ==================== Storage服务接口 ====================
service StorageService {
    // 健康检查
    common.HealthCheckResponse HealthCheck()
    
    // 文件上传
    GetUploadUrlResponse GetUploadUrl(1: GetUploadUrlRequest req)
    ConfirmUploadResponse ConfirmUpload(1: ConfirmUploadRequest req)
    
    // 文件下载和查询
    GetDownloadUrlResponse GetDownloadUrl(1: GetDownloadUrlRequest req)
    GetFileInfoResponse GetFileInfo(1: GetFileInfoRequest req)
    GetUserFilesResponse GetUserFiles(1: GetUserFilesRequest req)
    
    // 文件删除
    DeleteFileResponse DeleteFile(1: DeleteFileRequest req)
    BatchDeleteFilesResponse BatchDeleteFiles(1: BatchDeleteFilesRequest req)
}