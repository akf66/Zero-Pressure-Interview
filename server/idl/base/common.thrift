namespace go base

// 通用响应结构
struct BaseResponse {
    1: i64 status_code   // 0-成功，其他-失败
    2: string status_msg
}

// 空响应（用于HTTP层简单接口）
struct NilResponse {}

// 分页请求
struct PageRequest {
    1: i32 page = 1       // 页码，从1开始
    2: i32 page_size = 10 // 每页数量
}

// 分页响应
struct PageResponse {
    1: i64 total          // 总数
    2: i32 page           // 当前页
    3: i32 page_size      // 每页数量
}