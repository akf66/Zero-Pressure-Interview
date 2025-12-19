namespace go storage

include "../base/common.thrift"

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

// ==================== 文件服务接口 ====================
service StorageService {
    GetUploadUrlResponse GetUploadUrl(1: GetUploadUrlRequest req) (api.get="/api/v1/file/upload-url"),
    ConfirmUploadResponse ConfirmUpload(1: ConfirmUploadRequest req) (api.post="/api/v1/file/confirm"),
}