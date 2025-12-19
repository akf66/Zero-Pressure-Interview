namespace go base

// 文件信息
struct FileInfo {
    1: string file_key
    2: string file_name
    3: string content_type
    4: i64 file_size          // 文件大小（字节）
    5: i64 created_at
}

// 文件类型枚举
enum FileType {
    FT_NOT_SPECIFIED = 0  // 未指定
    RESUME = 1            // 简历
    AVATAR = 2            // 头像
    RECORDING = 3         // 录音
}