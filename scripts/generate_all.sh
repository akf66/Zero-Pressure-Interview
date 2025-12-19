#!/bin/bash

# Zero-Pressure-Interview 代码生成脚本
# 用于生成所有 Kitex 和 Hertz 代码

set -e

echo "=========================================="
echo "开始生成所有代码..."
echo "=========================================="

# 获取项目根目录
PROJECT_ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$PROJECT_ROOT"

# 1. 生成共享代码
echo ""
echo "步骤 1/3: 生成共享 Kitex 代码..."
bash scripts/generate_shared.sh

# 2. 生成 RPC 服务代码
echo ""
echo "步骤 2/3: 生成 RPC 服务代码..."
bash scripts/generate_rpc.sh

# 3. 生成 HTTP Gateway 代码
echo ""
echo "步骤 3/3: 生成 HTTP Gateway 代码..."
bash scripts/generate_http.sh

echo ""
echo "=========================================="
echo "所有代码生成完成！"
echo "=========================================="