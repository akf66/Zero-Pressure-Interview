#!/bin/bash

# 生成 RPC 服务代码

set -e

echo "生成 RPC 服务代码..."

# 获取项目根目录
PROJECT_ROOT=$(cd "$(dirname "$0")/.." && pwd)

# User 服务
echo ""
echo "生成 User 服务代码..."
cd "$PROJECT_ROOT/server/cmd/user"
kitex -service user \
      -module zpi \
      -use zpi/server/shared/kitex_gen \
      ../../idl/rpc/user.thrift

# Agent 服务
echo ""
echo "生成 Agent 服务代码..."
mkdir -p "$PROJECT_ROOT/server/cmd/agent"
cd "$PROJECT_ROOT/server/cmd/agent"
kitex -service agent \
      -module zpi \
      -use zpi/server/shared/kitex_gen \
      ../../idl/rpc/agent.thrift

# Question 服务
echo ""
echo "生成 Question 服务代码..."
mkdir -p "$PROJECT_ROOT/server/cmd/question"
cd "$PROJECT_ROOT/server/cmd/question"
kitex -service question \
      -module zpi \
      -use zpi/server/shared/kitex_gen \
      ../../idl/rpc/question.thrift

# Storage 服务
echo ""
echo "生成 Storage 服务代码..."
mkdir -p "$PROJECT_ROOT/server/cmd/storage"
cd "$PROJECT_ROOT/server/cmd/storage"
kitex -service storage \
      -module zpi \
      -use zpi/server/shared/kitex_gen \
      ../../idl/rpc/storage.thrift

echo ""
echo "RPC 服务代码生成完成！"