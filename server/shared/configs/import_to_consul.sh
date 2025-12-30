#!/bin/bash

# Consul KV 配置导入脚本
# 使用方法: ./import_to_consul.sh [consul_addr]
# 示例: ./import_to_consul.sh localhost:8500

CONSUL_ADDR=${1:-"localhost:8500"}
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "=== Zero Pressure Interview - Consul 配置导入 ==="
echo "Consul 地址: $CONSUL_ADDR"
echo ""

# 检查 Consul 是否可用
if ! curl -s "http://$CONSUL_ADDR/v1/status/leader" > /dev/null; then
    echo "❌ 无法连接到 Consul: $CONSUL_ADDR"
    exit 1
fi

echo "✅ Consul 连接成功"
echo ""

# 导入配置函数
import_config() {
    local key=$1
    local file=$2
    
    if [ -f "$SCRIPT_DIR/$file" ]; then
        echo -n "导入 $key ... "
        if curl -s --request PUT \
            --data @"$SCRIPT_DIR/$file" \
            "http://$CONSUL_ADDR/v1/kv/$key" > /dev/null; then
            echo "✅"
        else
            echo "❌"
        fi
    else
        echo "⚠️  文件不存在: $file"
    fi
}

# 导入各服务配置
import_config "zpi/api_srv" "api_srv.json"
import_config "zpi/user_srv" "user_srv.json"
import_config "zpi/interview_srv" "interview_srv.json"
import_config "zpi/question_srv" "question_srv.json"
import_config "zpi/storage_srv" "storage_srv.json"

echo ""
echo "=== 导入完成 ==="
echo ""
echo "验证配置:"
echo "  curl http://$CONSUL_ADDR/v1/kv/zpi/?keys"
