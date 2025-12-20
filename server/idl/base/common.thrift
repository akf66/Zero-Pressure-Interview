namespace go base

include "base.thrift"

// 通用响应结构
struct BaseResponse {
    1: i64 status_code        // 错误码，使用 base.ErrCode 枚举
    2: string status_msg      // 错误消息
    3: optional string error_source  // 错误来源服务（用于服务间调用错误传播）
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

// 健康检查响应
struct HealthCheckResponse {
    1: string status      // "healthy" / "unhealthy" / "degraded"
    2: string version     // 服务版本
    3: i64 uptime         // 运行时间（秒）
    4: map<string, string> details  // 详细信息（如数据库连接状态等）
}