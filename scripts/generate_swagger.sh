#!/bin/bash

# 一键生成所有服务的 Swagger 文档
# 用法: ./scripts/generate_swagger.sh

set -e  # 遇到错误立即退出

echo "========================================="
echo "开始生成 Swagger 文档..."
echo "========================================="

# 项目根目录
PROJECT_ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$PROJECT_ROOT"

# Swagger 输出目录
SWAGGER_DIR="swagger"

# 清理旧的 swagger 目录
if [ -d "$SWAGGER_DIR" ]; then
    echo "清理旧的 Swagger 文档..."
    rm -rf "$SWAGGER_DIR"
fi

# 创建 swagger 目录
mkdir -p "$SWAGGER_DIR"

echo ""
echo "开始生成各服务的 Swagger 文档..."
echo "-----------------------------------------"

# 遍历所有 HTTP IDL 文件
for file in ./server/idl/http/*.thrift; do
    # 获取文件名（不含扩展名）
    name=$(basename "$file" .thrift)
    
    echo "  - 生成 $name 服务的 Swagger 文档..."
    
    # 为每个服务创建独立目录
    output_dir="$SWAGGER_DIR/$name"
    mkdir -p "$output_dir"
    
    # 获取 IDL 文件的绝对路径
    idl_file="$PROJECT_ROOT/$file"
    
    # 切换到输出目录，在该目录下生成文档
    (
        cd "$output_dir"
        thriftgo -g go -p http-swagger "$idl_file"
    )
    
    if [ $? -eq 0 ]; then
        echo "    ✓ $name 服务 Swagger 文档生成成功: $output_dir/openapi.yaml"
    else
        echo "    ✗ $name 服务 Swagger 文档生成失败"
        exit 1
    fi
done

echo ""
echo "========================================="
echo "Swagger 文档生成完成！"
echo "========================================="
echo ""
echo "生成的文档位置："
echo "  - User 服务:      $SWAGGER_DIR/user/openapi.yaml"
echo "  - Interview 服务: $SWAGGER_DIR/interview/openapi.yaml"
echo "  - Question 服务:  $SWAGGER_DIR/question/openapi.yaml"
echo "  - Storage 服务:   $SWAGGER_DIR/storage/openapi.yaml"
echo ""
echo "下一步操作："
echo "1. 可以使用 Swagger UI 查看文档"
echo "2. 或者使用在线工具: https://editor.swagger.io/"
echo "3. 或者安装 swagger-ui: npm install -g swagger-ui-watcher"
echo "   然后运行: swagger-ui-watcher $SWAGGER_DIR/user/openapi.yaml"
echo ""