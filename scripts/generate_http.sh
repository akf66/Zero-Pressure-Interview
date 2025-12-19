#!/bin/bash

# 生成 HTTP Gateway 代码

set -e

echo "生成 HTTP Gateway 代码..."

# 获取项目根目录
PROJECT_ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$PROJECT_ROOT/server/cmd/api"

# 检查是否已存在 API 代码
if [ -f "main.go" ]; then
    echo "检测到已存在的 API 代码，执行更新..."
    hz update -idl ../../idl/http/api.thrift
else
    echo "首次生成 API 代码..."
    hz new -idl ../../idl/http/api.thrift -module zpi/server/cmd/api
fi

echo ""
echo "HTTP Gateway 代码生成完成！"