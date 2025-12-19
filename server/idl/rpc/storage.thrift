namespace go storage

include "../base/base.thrift"

// ==================== 获取上传URL ====================
struct GetUploadUrlReq {
    1: required i64 user_id
    2: required string file_name      // 文件名
    3: required string file_type      // 文件类型：resume/avatar/recording
    4: optional string content_type   // MIME类型，如：application/pdf
}

struct GetUploadUrlResp {
    1: base.BaseResp base_resp
    2: optional string upload_url     // 预签名上传URL
    3: optional string file_key       // 文件存储key
    4: optional i64 expires_at        // 过期时间戳（秒）
}

// ==================== 确认上传 ====================
struct ConfirmUploadReq {
    1: required i64 user_id
    2: required string file_key       // 文件存储key3: required string file_type      // 文件类型
}

struct ConfirmUploadResp {
    1: base.BaseResp base_resp
    2: optional string file_url       // 文件访问URL
}

// ==================== 获取下载URL ====================
struct GetDownloadUrlReq {
    1: required i64 user_id
    2: required string file_key       // 文件存储key
}

struct GetDownloadUrlResp {
    1: base.BaseResp base_resp
    2: optional string download_url   // 预签名下载URL
    3: optional i64 expires_at        // 过期时间戳（秒）
}

// ==================== 删除文件 ====================
struct DeleteFileReq {
    1: required i64 user_id
    2: required string file_key       // 文件存储key
}

struct DeleteFileResp {
    1: base.BaseResp base_resp
}

// ==================== 获取文件信息 ====================
struct GetFileInfoReq {
    1: required string file_key       // 文件存储key
}

struct FileInfo {
    1: string file_key
    2: string file_name
    3: string content_type
    4: i64 file_size                  // 文件大小（字节）
    5: i64 created_at
}

struct GetFileInfoResp {
    1: base.BaseResp base_resp
    2: optional FileInfo file_info
}

// ==================== 批量删除文件 ====================
struct BatchDeleteFilesReq {
    1: required i64 user_id
    2: required list<string> file_keys
}

struct BatchDeleteFilesResp {
    1: base.BaseResp base_resp
    2: optional i32 deleted_count     // 成功删除的数量
}

// ==================== 获取用户文件列表 ====================
struct GetUserFilesReq {
    1: required i64 user_id
    2: optional string file_type      // 文件类型筛选
    3: optional base.PageReq page
}

struct GetUserFilesResp {
    1: base.BaseResp base_resp
    2: optional list<FileInfo> files
    3: optional base.PageResp page
}

// ==================== Storage服务接口 ====================
service StorageService {
    // 获取上传URL
    GetUploadUrlResp GetUploadUrl(1: GetUploadUrlReq req)
    
    // 确认上传
    ConfirmUploadResp ConfirmUpload(1: ConfirmUploadReq req)
    
    // 获取下载URL
    GetDownloadUrlResp GetDownloadUrl(1: GetDownloadUrlReq req)
    
    // 删除文件
    DeleteFileResp DeleteFile(1: DeleteFileReq req)
    
    // 获取文件信息
    GetFileInfoResp GetFileInfo(1: GetFileInfoReq req)
    
    // 批量删除文件
    BatchDeleteFilesResp BatchDeleteFiles(1: BatchDeleteFilesReq req)
    
    // 获取用户文件列表
    GetUserFilesResp GetUserFiles(1: GetUserFilesReq req)
}