# Zero-Pressure-Interview (零压面试)

一个帮助求职者进行面试练习的AI智能面试平台，通过模拟真实面试场景，降低求职者的面试焦虑，提升面试能力。

## 项目特点

- 🤖 **AI智能面试官** - 基于大模型的智能对话，模拟真实面试场景
- 📝 **专项训练** - 针对特定技术栈的深度训练（Golang、MySQL、Redis等）
- 🎯 **综合面试** - 模拟完整的企业面试流程（一面、二面、三面、HR面）
- 📊 **能力分析** - 多维度能力评估，生成雷达图和详细报告
- 📚 **题库系统** - 丰富的面试题库，支持分类、搜索、收藏
- 📄 **简历分析** - AI分析简历，提供优化建议和岗位匹配

## 技术架构

### 微服务架构

```
API Gateway (Hertz) → RPC Services (Kitex)
                      ├── User Service (用户服务)
                      ├── Agent Service (AI面试服务)
                      ├── Question Service (题库服务)
                      └── Storage Service (文件存储服务)
```

### 技术栈

| 层级 | 技术 |
|------|------|
| HTTP框架 | Hertz |
| RPC框架 | Kitex |
| IDL | Thrift |
| 数据库 | MySQL 8.0 |
| 缓存 | Redis 7.0 |
| 对象存储 | MinIO |
| AI框架 | Eino ADK |
| 服务注册 | Etcd |

## 项目结构

```
Zero-Pressure-Interview/
├── docs/                      # 项目文档
│   ├── 需求文档.md
│   ├── 技术栈.md
│   ├── 执行目录.md
│   └── 项目结构说明.md
├── scripts/                   # 自动化脚本
│   ├── generate_all.sh       # 生成所有代码
│   ├── generate_shared.sh    # 生成共享代码
│   ├── generate_rpc.sh       # 生成RPC服务
│   └── generate_http.sh      # 生成HTTP Gateway
└── server/
    ├── cmd/                   # 服务入口
    │   ├── api/              # HTTP Gateway
    │   ├── user/             # User RPC服务
    │   ├── agent/            # Agent RPC服务
    │   ├── question/         # Question RPC服务
    │   └── storage/          # Storage RPC服务
    ├── idl/                   # 接口定义
    │   ├── base/             # 基础类型
    │   ├── rpc/              # RPC接口
    │   └── http/             # HTTP接口
    └── shared/                # 共享代码
        ├── consts/
        └── kitex_gen/
```

## 快速开始

### 1. 安装依赖工具

```bash
# 安装 Kitex 工具
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

# 安装 Thriftgo
go install github.com/cloudwego/thriftgo@latest

# 安装 Hertz 工具
go install github.com/cloudwego/hertz/cmd/hz@latest
```

### 2. 生成代码

```bash
# 生成所有服务代码
./scripts/generate_all.sh

# 或分步生成
./scripts/generate_shared.sh    # 生成共享代码
./scripts/generate_rpc.sh       # 生成RPC服务
./scripts/generate_http.sh      # 生成HTTP Gateway
```

### 3. 启动基础设施

使用 Docker Compose 启动所需的基础服务：

```bash
# 启动 MySQL、Redis、MinIO、Etcd
docker-compose up -d
```

### 4. 运行服务

```bash
# 启动 User 服务
cd server/cmd/user && go run .

# 启动 Agent 服务
cd server/cmd/agent && go run .

# 启动 Question 服务
cd server/cmd/question && go run .

# 启动 Storage 服务
cd server/cmd/storage && go run .

# 启动 API Gateway
cd server/cmd/api && go run .
```

## IDL 文件

项目使用 Thrift IDL 定义所有接口：

- [`base.thrift`](server/idl/base/base.thrift) - 基础类型定义
- [`user.thrift`](server/idl/rpc/user.thrift) - 用户服务接口
- [`agent.thrift`](server/idl/rpc/agent.thrift) - AI面试服务接口
- [`question.thrift`](server/idl/rpc/question.thrift) - 题库服务接口
- [`storage.thrift`](server/idl/rpc/storage.thrift) - 文件存储服务接口
- [`api.thrift`](server/idl/http/api.thrift) - HTTP Gateway接口

详细说明请查看 [IDL使用说明](server/idl/README.md)

## API 文档

### 用户接口

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/v1/user/register | 用户注册 |
| POST | /api/v1/user/login | 用户登录 |
| GET | /api/v1/user/profile | 获取个人信息 |
| PUT | /api/v1/user/profile | 更新个人信息 |
| POST | /api/v1/user/resume | 上传简历 |

### 面试接口

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/v1/interview/start | 开始面试 |
| POST | /api/v1/interview/:id/answer | 提交回答 |
| POST | /api/v1/interview/:id/finish | 结束面试 |
| GET | /api/v1/interview/history | 面试历史 |
| GET | /api/v1/interview/:id | 面试详情 |

### 题库接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /api/v1/question/categories | 获取分类列表 |
| GET | /api/v1/question/list | 获取题目列表 |
| GET | /api/v1/question/:id | 获取题目详情 |
| POST | /api/v1/question/favorite | 收藏题目 |

## 文档

- [需求文档](docs/需求文档.md) - 详细的功能需求和数据模型
- [技术栈文档](docs/技术栈.md) - 技术选型和使用说明
- [项目结构说明](docs/项目结构说明.md) - 完整的项目结构说明
- [执行目录](docs/执行目录.md) - 代码生成命令参考

## 开发计划

- [x] 项目初始化和目录结构
- [x] IDL 接口定义
- [ ] 生成服务代码
- [ ] 实现用户服务
- [ ] 实现题库服务
- [ ] 实现AI面试服务
- [ ] 实现文件存储服务
- [ ] 实现API Gateway
- [ ] 集成AI模型
- [ ] 前端开发

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 联系方式

如有问题，请提交 Issue 或联系项目维护者。
