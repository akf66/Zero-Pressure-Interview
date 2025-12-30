# Consul KV 配置文件

本目录包含各微服务在 Consul 中的 KV 配置模板。

## 使用方法

### 1. 通过 Consul UI 导入
1. 访问 Consul UI: http://localhost:8500/ui
2. 进入 Key/Value 页面
3. 创建对应的 Key，将 JSON 内容粘贴进去

### 2. 通过 Consul CLI 导入
```bash
# 导入 API 网关配置
consul kv put zpi/api_srv @api_srv.json

# 导入 User 服务配置
consul kv put zpi/user_srv @user_srv.json

# 导入 Interview 服务配置
consul kv put zpi/interview_srv @interview_srv.json

# 导入 Question 服务配置
consul kv put zpi/question_srv @question_srv.json

# 导入 Storage 服务配置
consul kv put zpi/storage_srv @storage_srv.json
```

### 3. 通过 HTTP API 导入
```bash
# 示例：导入 User 服务配置
curl --request PUT \
  --data @user_srv.json \
  http://localhost:8500/v1/kv/zpi/user_srv
```

## 配置文件说明

| 文件 | Consul Key | 服务 |
|------|------------|------|
| api_srv.json | zpi/api_srv | API 网关 |
| user_srv.json | zpi/user_srv | 用户服务 |
| interview_srv.json | zpi/interview_srv | 面试服务 |
| question_srv.json | zpi/question_srv | 题库服务 |
| storage_srv.json | zpi/storage_srv | 存储服务 |

## 注意事项

1. **敏感信息**: 生产环境请修改所有密码、密钥等敏感信息
2. **端口冲突**: 确保各服务端口不冲突
3. **数据库**: 确保 MySQL、Redis 服务已启动
4. **服务发现**: 确保 Consul 服务已启动
