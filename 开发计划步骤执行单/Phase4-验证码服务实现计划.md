# Phase 4: 验证码服务实现计划

## 目标
实现完整的验证码服务，包括邮件验证码、短信验证码、Redis 存储和 Token 黑名单机制。

## 实现步骤

### 1. Redis 集成
- [ ] 添加 Redis 配置到 config.yaml
- [ ] 实现 Redis 客户端初始化
- [ ] 创建 Redis 工具包

### 2. 验证码生成器
- [ ] 创建验证码生成工具（6位数字）
- [ ] 实现验证码存储到 Redis（5分钟过期）
- [ ] 实现验证码验证逻辑

### 3. 邮件服务
- [ ] 添加邮件配置（SMTP）
- [ ] 实现邮件发送工具
- [ ] 创建验证码邮件模板

### 4. 短信服务（可选）
- [ ] 集成短信服务商 API（阿里云/腾讯云）
- [ ] 实现短信发送工具
- [ ] 创建短信模板

### 5. 实现 SendVerifyCode
- [ ] 参数验证
- [ ] 生成验证码
- [ ] 根据类型发送邮件/短信
- [ ] 存储到 Redis

### 6. 实现 ResetPassword
- [ ] 验证验证码
- [ ] 查找用户
- [ ] 更新密码
- [ ] 删除已使用的验证码

### 7. 完善 Register 和 Login
- [ ] Register: 验证验证码
- [ ] Login: 支持验证码登录

### 8. Token 黑名单
- [ ] 实现 Token 黑名单存储
- [ ] 完善 Logout 功能
- [ ] 创建 Token 验证中间件

## 技术选型

### Redis
- **库**: `github.com/redis/go-redis/v9`
- **用途**: 
  - 验证码存储（key: `verify_code:{target}`, TTL: 5分钟）
  - Token 黑名单（key: `token_blacklist:{token}`, TTL: Token剩余时间）
  - 限流控制（防止验证码滥发）

### 邮件服务
- **库**: `gopkg.in/gomail.v2`
- **配置**: SMTP 服务器（支持 Gmail, QQ邮箱等）

### 短信服务（可选）
- **阿里云**: `github.com/aliyun/alibaba-cloud-sdk-go`
- **腾讯云**: `github.com/tencentcloud/tencentcloud-sdk-go`

## 配置示例

```yaml
# config.yaml
redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
  pool_size: 10

email:
  smtp_host: smtp.gmail.com
  smtp_port: 587
  username: your-email@gmail.com
  password: your-app-password
  from_name: "Zero Pressure Interview"

sms:
  provider: aliyun  # aliyun, tencent
  access_key_id: your-access-key
  access_key_secret: your-secret
  sign_name: "零压面试"
  template_code: "SMS_123456"
```

## 数据结构

### Redis Key 设计
```
# 验证码
verify_code:email:{email}:{purpose} = "123456"  # TTL: 5分钟
verify_code:phone:{phone}:{purpose} = "123456"  # TTL: 5分钟

# 限流
rate_limit:verify_code:{target} = count  # TTL: 1小时

# Token 黑名单
token_blacklist:{token_hash} = "1"  # TTL: Token剩余时间
```

### 验证码用途枚举
- `REGISTER` = 1: 注册
- `LOGIN` = 2: 登录
- `RESET_PASSWORD` = 3: 重置密码
- `CHANGE_PHONE` = 4: 修改手机号
- `CHANGE_EMAIL` = 5: 修改邮箱

## 安全考虑

1. **限流**: 同一目标1小时内最多发送5次
2. **验证码复杂度**: 6位随机数字
3. **有效期**: 5分钟
4. **一次性使用**: 验证后立即删除
5. **防暴力破解**: 验证失败3次后锁定10分钟

## 测试计划

1. 单元测试- 验证码生成
   - Redis 存储和读取
   - 邮件发送（Mock）

2. 集成测试
   - 完整的注册流程
   - 密码重置流程
   - 验证码登录流程

3. 性能测试
   - Redis 连接池
   - 并发验证码发送

## 预计工作量
- Redis 集成: 1小时
- 验证码生成器: 30分钟
- 邮件服务: 1小时
- 实现业务逻辑: 2小时
- 测试和调试: 1小时
- **总计**: 约 5.5 小时