# IDL 设计指南

> 基于 FreeCar 项目总结的 Thrift IDL 设计最佳实践

## 📁 目录结构

```
your-project/
├── idl/
│   ├── base/                    # 基础层：共享数据模型
│   │   ├── common.thrift        # 通用响应结构（BaseResponse、NilResponse）
│   │   ├── user.thrift# 用户领域模型（User、UserEntity、UserStatus）
│   │   ├── order.thrift         # 订单领域模型
│   │   └── ...  # 其他领域模型
│   │
│   ├── http/                    # HTTP 层：API 网关接口（Hertz）
│   │   ├── user.thrift          # 用户 HTTP API（带 api.post/get 注解）
│   │   ├── order.thrift         # 订单 HTTP API
│   │   └── ...
│   │
│   └── rpc/                     # RPC 层：微服务接口（Kitex）
│       ├── user.thrift          # 用户 RPC 服务（完整 Request/Response）
│       ├── order.thrift         # 订单 RPC 服务
│       └── ...
│
└── server/
    └── cmd/├── api/                 # HTTP 网关服务（调用 RPC）├── user/                # 用户微服务
        ├── order/               # 订单微服务
        └── ...
```

**三层职责说明**：

| 层级 | 职责 | 特点 | 被谁使用 |
|------|------|------|----------|
| **base/** | 定义数据结构、枚举 | 纯数据定义，无业务逻辑 | http/ 和 rpc/ 通过 `include` 引用 |
| **http/** | 定义 RESTful API | 包含 HTTP 注解（路由、校验） | 客户端（Web/App）调用 |
| **rpc/** | 定义服务间接口 | 完整的 Request/Response | 微服务之间调用 |
```

## 🎯 核心原则

1. **三层分离**：Base 定义数据，HTTP 定义 API，RPC 定义服务
2. **DRY**：通过 `include` 复用 Base 层定义
3. **向后兼容**：字段编号递增，不删除不重用

## 📝 命名规范

| 类型 | 规则 | 示例 |
|------|------|------|
| 请求 | `{Action}{Resource}Request` | `CreateUserRequest` |
| 响应 | `{Action}{Resource}Response` | `CreateUserResponse` |
| 实体 | `{Resource}Entity` | `UserEntity` |
| 服务 | `{Resource}Service` | `UserService` |
| 方法 | `{Action}{Resource}` | `CreateUser`, `GetUser` |
| 管理员方法 | `Admin{Action}{Resource}` | `AdminDeleteUser` |

## 💻 代码模板

### Base 层 - common.thrift

```thrift
namespace go base

struct BaseResponse {
    1: i64 status_code,   // 0-成功，其他-失败
    2: string status_msg,
}

struct NilResponse {}
```

### Base 层 - 领域模型

```thrift
namespace go base

// 实体（带 ID）
struct UserEntity {
    1: string id,
    2: User user,
}

// 领域对象
struct User {
    1: string username,
    2: string email,
    3: i64 created_at,
}

// 枚举（必须有 0 值）
enum UserStatus {
    US_NOT_SPECIFIED = 0,
    ACTIVE = 1,
    INACTIVE = 2,
}
```

### HTTP 层

```thrift
namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

struct CreateUserRequest {
    1: string username (api.raw = "username", api.vd = "len($) > 0 && len($) < 50"),
    2: string email (api.raw = "email"),
}

service UserService {
    // 管理后台
    common.NilResponse AdminDeleteUser(1: DeleteUserRequest req) (api.delete = "/admin/user"),
    
    // 客户端
    common.NilResponse CreateUser(1: CreateUserRequest req) (api.post = "/user"),
    common.NilResponse GetUser(1: GetUserRequest req) (api.get = "/user"),
}
```

**常用注解**：
- `api.raw = "field"` - 参数绑定
- `api.vd = "len($) > 0"` - 长度校验
- `api.vd = "$ > 0 && $ < 100"` - 数值范围
- `api.get/post/put/delete = "/path"` - HTTP 路由

### RPC 层

```thrift
namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

struct CreateUserRequest {
    1: string account_id,  // 调用者 ID
    2: string username,
    3: string email,
}

struct CreateUserResponse {
    1: common.BaseResponse base_resp,  // 必须包含
    2: user.UserEntity user_entity,
}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req),
    GetUserResponse GetUser(1: GetUserRequest req),
}
```

## ✅ 最佳实践

### 1. 字段编号

```thrift
struct User {
    1: string id,
    2: string username,
    // 3: string phone,  // 已废弃，保留编号
    4: string email,     // 新字段用下一个编号
}
```

### 2. 枚举默认值

```thrift
enum Status {
    NOT_SPECIFIED = 0,  // ✅ 必须有
    ACTIVE = 1,
    INACTIVE = 2,
}
```

### 3. 响应结构

```thrift
// RPC 层：完整响应
struct GetUserResponse {
    1: common.BaseResponse base_resp,  // ✅ 必须
    2: User user,
}

// HTTP 层：简单接口用 NilResponse
service UserService {
    common.NilResponse CreateUser(...) (api.post = "/user"),
}
```

## 🔧 代码生成

```bash
# HTTP 服务（Hertz）
hz new -module your-module -idl idl/http/user.thrift
hz update -idl idl/http/user.thrift

# RPC 服务（Kitex）
kitex -module your-module -service user idl/rpc/user.thrift
```

## 📚 参考

- [Thrift 文档](https://thrift.apache.org/docs/)
- [Hertz 文档](https://www.cloudwego.io/zh/docs/hertz/)
- [Kitex 文档](https://www.cloudwego.io/zh/docs/kitex/)