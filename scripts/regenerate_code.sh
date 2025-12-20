#!/bin/bash

# 重新生成所有 Kitex 和 Hertz 代码的脚本
# 用于应用 IDL 修改后重新生成代码

set -e  # 遇到错误立即退出

echo "========================================="
echo "开始重新生成代码..."
echo "========================================="

# 项目根目录
PROJECT_ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$PROJECT_ROOT"

echo ""
echo "1. 重新生成共享 Model 库..."
echo "-----------------------------------------"
cd server/shared

echo "  - 生成 user 相关代码..."
kitex -module zpi ../idl/rpc/user.thrift

echo "  - 生成 interview 相关代码..."
kitex -module zpi ../idl/rpc/interview.thrift

echo "  - 生成 question 相关代码..."
kitex -module zpi ../idl/rpc/question.thrift

echo "  - 生成 storage 相关代码..."
kitex -module zpi ../idl/rpc/storage.thrift

echo ""
echo "2. 更新 User RPC 服务..."
echo "-----------------------------------------"
cd ../cmd/user
kitex -service user \-module zpi \
      -use zpi/server/shared/kitex_gen \
      ../../idl/rpc/user.thrift

echo ""
echo "3. 更新 Interview RPC 服务..."
echo "-----------------------------------------"
cd ../interview
kitex -service interview \
      -module zpi \
      -use zpi/server/shared/kitex_gen \
      ../../idl/rpc/interview.thrift

echo ""
echo "4. 更新 Question RPC 服务..."
echo "-----------------------------------------"
cd ../question
kitex -service question \
      -module zpi \
      -use zpi/server/shared/kitex_gen \
      ../../idl/rpc/question.thrift

echo ""
echo "5. 更新 Storage RPC 服务..."
echo "-----------------------------------------"
cd ../storage
kitex -service storage \
      -module zpi \
      -use zpi/server/shared/kitex_gen \
      ../../idl/rpc/storage.thrift

echo ""
echo "6. 更新 HTTP 网关..."
echo "-----------------------------------------"
cd ../api

echo "  - 更新 user HTTP 接口..."
hz update -idl ../../idl/http/user.thrift

echo "  - 更新 interview HTTP 接口..."
hz update -idl ../../idl/http/interview.thrift

echo "  - 更新 question HTTP 接口..."
hz update -idl ../../idl/http/question.thrift

echo "  - 更新 storage HTTP 接口..."
hz update -idl ../../idl/http/storage.thrift

echo ""
echo "========================================="
echo "代码生成完成！"
echo "========================================="
echo ""
echo "下一步操作："
echo "1. 更新 server/shared/errno/errno.go 使用新的错误码"
echo "2. 为每个服务实现 HealthCheck() 方法"
echo "3. 运行 'go mod tidy' 更新依赖"
echo "4. 编译并测试所有服务"
echo ""