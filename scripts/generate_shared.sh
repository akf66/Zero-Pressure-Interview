#!/bin/bash

# 生成共享的 Kitex 代码

set -e

echo "生成共享 Kitex 代码..."

# 获取项目根目录
PROJECT_ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$PROJECT_ROOT/server/shared"

# 清理旧的生成代码
echo "清理旧的生成代码..."
rm -rf kitex_gen

# 生成各个服务的共享代码
echo "生成 User 服务共享代码..."
kitex -module zpi ../idl/rpc/user.thrift

echo "生成 Agent 服务共享代码..."
kitex -module zpi ../idl/rpc/agent.thrift

echo "生成 Question 服务共享代码..."
kitex -module zpi ../idl/rpc/question.thrift

echo "生成 Storage 服务共享代码..."
kitex -module zpi ../idl/rpc/storage.thrift

echo "共享代码生成完成！"