# IDL 架构说明

本项目采用三层分离的 Thrift IDL 架构设计，遵循最佳实践。

## 📁 目录结构

```
server/idl/
├── base/                    # 基础层：共享数据模型
│   ├── common.thrift        # 通用响应结构（BaseResponse、NilResponse、分页）
│   ├── user.thrift          # 用户领域模型（UserEntity、User、Resume）
│   ├── question.thrift      # 题目领域模型（QuestionEntity、Question）
│   ├── interview.thrift     # 面试领域模型（InterviewEntity、Interview、Message）
│   ├── storage.thrift       # 存储领域模型（FileInfo）
│   └── base.thrift          # 错误码常量定义
│
├── http/                    # HTTP 层：API 网关接口（Hertz）
│   ├── user.thrift          # 用户 HTTP API
│   ├── interview.thrift     # 面试 HTTP API
│   ├── question.thrift      # 题库 HTTP API
│   └── storage.thrift       # 文件 HTTP API
│
└── rpc/                     # RPC 层：微服务接口（Kitex）
    ├── user.thrift          # 用户 RPC 服务
    ├── agent.thrift         # AI面试 RPC 服务
    ├── question.thrift      # 题库 RPC 服务
    └── storage.thrift       # 存储 RPC 服务
```

## 🎯 三层职责

| 层级 | 职责 | 特点 | 被谁使用 |
|------|------|------|----------|
| **base/** | 定义数据结构、枚举 | 纯数据定义，无业务逻辑 | http/ 和 rpc/ 通过 `include` 引用 |
| **http/** | 定义 RESTful API | 包含 HTTP 注解（路由、校验） | 客户端（Web/App）调用 |
| **rpc/** | 定义服务间接口 | 完整的 Request/Response | 微服务之间调用 |

## 📝 命名规范

| 类型 | 规则 | 示例 |
|------|------|------|
| 请求 | `{Action}{Resource}Request` | `CreateUserRequest` |
| 响应 | `{Action}{Resource}Response` | `CreateUserResponse` |
| 实体 | `{Resource}Entity` | `UserEntity` |
| 领域对象 | `{Resource}` | `User` |
| 服务 | `{Resource}Service` | `UserService` |
| 方法 | `{Action}{Resource}` | `CreateUser`, `GetUser` |

## 🔄 代码生成

### HTTP 服务（Hertz）

```bash
# 生成用户服务
hz new -module github.com/your-org/zpi -idl server/idl/http/user.thrift

# 更新服务
hz update -idl server/idl/http/user.thrift
```

### RPC 服务（Kitex）

```bash
# 生成用户服务
kitex -module github.com/your-org/zpi -service user server/idl/rpc/user.thrift

# 生成共享代码
kitex -module github.com/your-org/zpi server/idl/base/common.thrift
```

## ✨ 设计亮点

1. **三层分离**：Base 定义数据，HTTP 定义 API，RPC 定义服务
2. **DRY 原则**：通过 `include` 复用 Base 层定义，避免重复
3. **领域驱动**：按业务领域拆分文件（user、question、interview、storage）
4. **类型安全**：使用 Entity 和领域对象分离，ID 和数据分开管理
5. **枚举规范**：所有枚举都有 0 值（NOT_SPECIFIED）
6. **向后兼容**：字段编号递增，不删除不重用

## 📚 参考文档

- [IDL架构设计文档](../../docs/idl架构设计.md)
- [Thrift 官方文档](https://thrift.apache.org/docs/)
- [Hertz 文档](https://www.cloudwego.io/zh/docs/hertz/)
- [Kitex 文档](https://www.cloudwego.io/zh/docs/kitex/)