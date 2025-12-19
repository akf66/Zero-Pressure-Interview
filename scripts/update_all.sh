#!/bin/bash

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_ROOT=$(cd "$(dirname "$0")/.." && pwd)

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}开始更新所有微服务代码${NC}"
echo -e "${GREEN}========================================${NC}"

# 1. 生成共享的 kitex_gen 代码（使用RPC层的thrift）
echo -e "\n${YELLOW}[1/6] 生成共享的 RPC 代码...${NC}"
cd "$PROJECT_ROOT/server/shared" || exit 1

echo "  - 生成 user.thrift..."
kitex -module zpi ../idl/rpc/user.thrift

echo "  - 生成 interview.thrift..."
kitex -module zpi ../idl/rpc/interview.thrift

echo "  - 生成 question.thrift..."
kitex -module zpi ../idl/rpc/question.thrift

echo "  - 生成 storage.thrift..."
kitex -module zpi ../idl/rpc/storage.thrift

echo -e "${GREEN}✓ 共享 RPC 代码生成完成${NC}"

# 2. 更新 HTTP API 网关
echo -e "\n${YELLOW}[2/6] 更新 HTTP API 网关...${NC}"
cd "$PROJECT_ROOT/server/cmd/api" || exit 1
make all
echo -e "${GREEN}✓ API 网关更新完成${NC}"

# 3. 更新 User 服务
echo -e "\n${YELLOW}[3/6] 更新 User 服务...${NC}"
cd "$PROJECT_ROOT/server/cmd/user" || exit 1
make server
echo -e "${GREEN}✓ User 服务更新完成${NC}"

# 4. 更新 Interview 服务
echo -e "\n${YELLOW}[4/6] 更新 Interview 服务...${NC}"
cd "$PROJECT_ROOT/server/cmd/interview" || exit 1
make server
echo -e "${GREEN}✓ Interview 服务更新完成${NC}"

# 5. 更新 Question 服务
echo -e "\n${YELLOW}[5/6] 更新 Question 服务...${NC}"
cd "$PROJECT_ROOT/server/cmd/question" || exit 1
make server
echo -e "${GREEN}✓ Question 服务更新完成${NC}"

# 6. 更新 Storage 服务
echo -e "\n${YELLOW}[6/6] 更新 Storage 服务...${NC}"
cd "$PROJECT_ROOT/server/cmd/storage" || exit 1
make server
echo -e "${GREEN}✓ Storage 服务更新完成${NC}"

# 完成
echo -e "\n${GREEN}========================================${NC}"
echo -e "${GREEN}所有微服务代码更新完成！${NC}"
echo -e "${GREEN}========================================${NC}"

cd "$PROJECT_ROOT" || exit 1