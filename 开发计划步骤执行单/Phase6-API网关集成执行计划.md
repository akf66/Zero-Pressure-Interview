# Phase 6: API 网关集成执行计划

## 目标
实现 API 网关与各个微服务的集成，包括 RPC 客户端调用、认证中间件、HTTP 路由配置等。

## 当前状态
- ✅ User 服务已完整实现（注册、登录、验证码、简历管理等）
- ✅ IDL 定义已完成
- ✅ 代码生成已完成
- ⏸️ API 网关需要集成 User 服务

## 实现步骤

### 1. 分析现有 API 网关结构
- [ ] 查看 `server/cmd/api` 目录结构
- [ ] 分析现有的 RPC 客户端初始化代码
- [ ] 查看现有的 handler 实现
- [ ] 了解路由配置方式

### 2. 实现 User 服务 RPC 客户端
- [ ] 检查 `initialize/rpc/user_service.go` 是否正确初始化
- [ ] 配置服务发现（Consul）
- [ ] 配置连接池和超时设置
- [ ] 实现健康检查

### 3. 实现认证中间件
- [ ] 创建 Token 验证中间件
- [ ] 实现 Paseto Token 解析
- [ ] 从 Token 中提取用户信息
- [ ] 将用户信息注入到上下文
- [ ] 处理 Token 过期和无效情况

### 4. 实现 User Handler
- [ ] 实现注册接口 `/api/v1/user/register`
- [ ] 实现登录接口 `/api/v1/user/login`
- [ ] 实现登出接口 `/api/v1/user/logout`
- [ ] 实现获取用户信息 `/api/v1/user/profile`
- [ ] 实现更新用户信息 `/api/v1/user/profile`
- [ ] 实现修改密码 `/api/v1/user/password`
- [ ] 实现重置密码 `/api/v1/user/password/reset`
- [ ] 实现发送验证码 `/api/v1/user/verify-code`
- [ ] 实现简历管理接口
- [ ] 实现注销账号 `/api/v1/user/account`

### 5. 配置路由和中间件
- [ ] 配置公开路由（不需要认证）
  - 注册
  - 登录
  - 发送验证码
  - 重置密码
- [ ] 配置需要认证的路由
  - 获取/更新用户信息
  - 修改密码
  - 简历管理
  - 登出
  - 注销账号
- [ ] 应用认证中间件到需要认证的路由组

### 6. 错误处理和响应转换
- [ ] 实现 RPC 错误到 HTTP 错误的转换
- [ ] 统一响应格式
- [ ] 实现错误日志记录

### 7. 测试和验证
- [ ] 测试注册流程
- [ ] 测试登录流程（密码/验证码）
- [ ] 测试 Token 认证
- [ ] 测试用户信息管理
- [ ] 测试简历管理
- [ ] 测试错误处理

## 技术要点

### RPC 客户端配置
```go
// 服务发现配置
resolver: consul
registry: consul://localhost:8500

// 连接池配置
max_idle_conns: 10
max_idle_per_host: 5

// 超时配置
connect_timeout: 5s
rpc_timeout: 10s
```

### 认证中间件设计
```go
func AuthMiddleware() app.HandlerFunc {
    return func(ctx context.Context, c *app.RequestContext) {
        // 1. 从 Header 获取 Token
        token := c.GetHeader("Authorization")
        
        // 2. 验证 Token
        claims, err := paseto.ParseToken(token)
        if err != nil {
            // Token 无效
            c.JSON(401, ...)
            c.Abort()
            return
        }
        
        // 3. 检查 Token 是否在黑名单（Redis）
        // TODO: Phase 4 完善
        
        // 4. 将用户信息注入上下文
        c.Set("user_id", claims.ID)
        c.Set("username", claims.Subject)
        
        c.Next()
    }
}
```

### Handler 实现模式
```go
func Register(ctx context.Context, c *app.RequestContext) {
    var req user.RegisterRequest
    
    // 1. 绑定请求参数
    if err := c.BindAndValidate(&req); err != nil {
        c.JSON(400, ...)
        return
    }
    
    // 2. 调用 RPC 服务
    resp, err := rpc.UserClient.Register(ctx, &req)
    if err != nil {
        // RPC 调用失败
        c.JSON(500, ...)
        return
    }
    
    // 3. 转换响应
    c.JSON(200, resp)
}
```

## 路由配置

### 公开路由（无需认证）
```
POST   /api/v1/user/register        - 用户注册
POST   /api/v1/user/login           - 用户登录
POST   /api/v1/user/verify-code     - 发送验证码
POST   /api/v1/user/password/reset  - 重置密码
```

### 需要认证的路由
```
GET    /api/v1/user/profile         - 获取用户信息
PUT    /api/v1/user/profile         - 更新用户信息
POST   /api/v1/user/password        - 修改密码
POST   /api/v1/user/logout          - 退出登录
DELETE /api/v1/user/account         - 注销账号

POST   /api/v1/user/resume          - 上传简历
GET    /api/v1/user/resume          - 获取简历
```

## 文件结构

```
server/cmd/api/
├── biz/
│   ├── handler/
│   │   └── user/
│   │       └── user_service.go      # User Handler 实现
│   ├── router/
│   │   └── user/
│   │       ├── user.go              # 路由注册
│   │       └── middleware.go        # 认证中间件
│   └── model/
│       └── user/
│           └── user.go              # 请求/响应模型
├── initialize/
│   ├── rpc/
│   │   └── user_service.go          # User RPC 客户端初始化
│   └── cors.go                      # CORS 配置
└── main.go
```

## 配置文件

### config.yaml
```yaml
# RPC 服务配置
rpc:
  user:
    name: "zpi/user_srv"
    timeout: 10s
    
# Consul 配置
consul:
  host: "127.0.0.1"
  port: 8500
  
# Paseto 配置
paseto:
  secret_key: "your-32-byte-secret-key-here-change-it"
  implicit: "zpi-api-gateway"
```

## 预计工作量
- RPC 客户端集成: 1小时
- 认证中间件实现: 1.5小时
- Handler 实现: 2小时
- 路由配置: 0.5小时
- 测试和调试: 1小时
- **总计**: 约 6 小时

## 依赖项
- ✅ User 服务已实现
- ✅ IDL 定义已完成
- ✅ Consul 服务发现（需要运行）
- ⏸️ Paseto Token 验证库

## 成功标准
1. 所有 User 相关接口可通过 HTTP 访问
2. 认证中间件正确验证 Token
3. 错误处理完善
4. 响应格式统一
5. 日志记录完整

## 注意事项
1. Token 验证需要与 User 服务使用相同的密钥
2. 需要处理 RPC 超时和重试
3. 需要实现优雅的错误处理
4. 考虑并发安全
5. 注意跨域配置

## 后续优化
- 实现请求限流
- 实现 API 文档（Swagger）
- 实现请求日志中间件
- 实现性能监控
- 实现熔断降级