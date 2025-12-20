#!/bin/bash

# ============================================
# TLS/SSL 证书生成脚本
# 用于开发和测试环境
# ============================================

set -e  # 遇到错误立即退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 证书配置
CERT_DIR="$(cd "$(dirname "$0")" && pwd)"
CERT_FILE="server.crt"
KEY_FILE="server.key"
DAYS=365
KEY_SIZE=2048

# 证书主题信息
COUNTRY="CN"
STATE="Beijing"
CITY="Beijing"
ORGANIZATION="ZPI"
ORGANIZATIONAL_UNIT="Development"
COMMON_NAME="localhost"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  ZPI TLS 证书生成工具${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# 检查 OpenSSL 是否安装
if ! command -v openssl &> /dev/null; then
    echo -e "${RED}❌ 错误: 未找到 openssl 命令${NC}"
    echo -e "${YELLOW}请先安装 OpenSSL:${NC}"
    echo -e "  Ubuntu/Debian: sudo apt-get install openssl"
    echo -e "  CentOS/RHEL:   sudo yum install openssl"
    echo -e "  macOS:         brew install openssl"
    exit 1
fi

echo -e "${GREEN}✓${NC} OpenSSL 版本: $(openssl version)"
echo ""

# 切换到证书目录
cd "$CERT_DIR"

# 检查是否已存在证书
if [ -f "$CERT_FILE" ] || [ -f "$KEY_FILE" ]; then
    echo -e "${YELLOW}⚠️  警告: 证书文件已存在${NC}"
    echo -e "  - $CERT_FILE"
    echo -e "  - $KEY_FILE"
    echo ""
    read -p "是否覆盖现有证书? (y/N): " -n 1 -r
    echo ""
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${BLUE}已取消操作${NC}"
        exit 0
    fi
    echo ""
fi

# 显示证书配置信息
echo -e "${BLUE}证书配置信息:${NC}"
echo -e "  国家 (C):           $COUNTRY"
echo -e "  省份 (ST):          $STATE"
echo -e "  城市 (L):           $CITY"
echo -e "  组织 (O):           $ORGANIZATION"
echo -e "  部门 (OU):          $ORGANIZATIONAL_UNIT"
echo -e "  通用名称 (CN):      $COMMON_NAME"
echo -e "  密钥长度:           $KEY_SIZE bits"
echo -e "  有效期:             $DAYS 天"
echo -e "  证书文件:           $CERT_FILE"
echo -e "  私钥文件:           $KEY_FILE"
echo ""

# 生成证书
echo -e "${YELLOW}正在生成证书...${NC}"

openssl req -x509 -newkey rsa:$KEY_SIZE -nodes \
    -keyout "$KEY_FILE" \
    -out "$CERT_FILE" \
    -days $DAYS \
    -subj "/C=$COUNTRY/ST=$STATE/L=$CITY/O=$ORGANIZATION/OU=$ORGANIZATIONAL_UNIT/CN=$COMMON_NAME" \
    2>/dev/null

# 检查生成结果
if [ -f "$CERT_FILE" ] && [ -f "$KEY_FILE" ]; then
    echo -e "${GREEN}✅ 证书生成成功！${NC}"
    echo ""
    # 显示文件信息
    echo -e "${BLUE}生成的文件:${NC}"
    ls -lh "$CERT_FILE" "$KEY_FILE" | awk '{print "  " $9 " (" $5 ")"}'
    echo ""
    
    # 显示证书详细信息
    echo -e "${BLUE}证书详细信息:${NC}"
    echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    openssl x509 -in "$CERT_FILE" -noout -text | grep -A 2 "Subject:"
    openssl x509 -in "$CERT_FILE" -noout -text | grep -A 2 "Validity"
    openssl x509 -in "$CERT_FILE" -noout -text | grep "Public-Key:"
    echo -e "${GREEN}━━━━━━━━━━━━━━━━${NC}"
    echo ""
    
    # 显示证书指纹
    echo -e "${BLUE}证书指纹 (SHA256):${NC}"
    openssl x509 -in "$CERT_FILE" -noout -fingerprint -sha256 | sed 's/^/  /'
    echo ""
    
    # 使用提示
    echo -e "${BLUE}使用说明:${NC}"
    echo -e "  1. 证书文件路径: ${GREEN}$CERT_DIR/$CERT_FILE${NC}"
    echo -e "  2. 私钥文件路径: ${GREEN}$CERT_DIR/$KEY_FILE${NC}"
    echo -e "  3. 在代码中使用:"
    echo -e "     ${YELLOW}tls.LoadX509KeyPair(\"$CERT_DIR/$CERT_FILE\", \"$CERT_DIR/$KEY_FILE\")${NC}"
    echo ""
    
    # 安全提示
    echo -e "${YELLOW}⚠️  安全提示:${NC}"
    echo -e "  - 这是自签名证书，仅用于开发/测试环境"
    echo -e "  - 浏览器会显示 '不安全' 警告（正常现象）"
    echo -e "  - 生产环境请使用受信任的 CA 证书（如 Let's Encrypt）"
    echo -e "  - 请勿将私钥文件 ($KEY_FILE) 提交到 Git 仓库"
    echo ""
    
    # 设置文件权限
    chmod 644 "$CERT_FILE"
    chmod 600 "$KEY_FILE"
    echo -e "${GREEN}✓${NC} 已设置文件权限 (证书: 644, 私钥: 600)"
    echo ""
    
    echo -e "${GREEN}========================================${NC}"
    echo -e "${GREEN}  证书生成完成！${NC}"
    echo -e "${GREEN}========================================${NC}"
else
    echo -e "${RED}❌ 证书生成失败${NC}"
    exit 1
fi