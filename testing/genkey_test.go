package genkey_test

import (
	"encoding/hex"
	"testing"

	"aidanwoods.dev/go-paseto"
)

func TestGenTokenKey(t *testing.T) {
	// 生成密钥对
	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public()

	// 转换为十六进制字符串
	secretKeyHex := hex.EncodeToString(secretKey.ExportBytes())
	publicKeyHex := hex.EncodeToString(publicKey.ExportBytes())

	t.Logf("secretKeyHex: %s\n", secretKeyHex)
	t.Logf("publicKeyHex: %s\n", publicKeyHex)

}

//53d2d2c41f42200f761c5773dc0ca2e0df11201786621b99eb7559cb4ef399583c99b5671b0d97ff8f54edce9d61aa9f8d09a3cd691706d4f661c912551f83bb
//3c99b5671b0d97ff8f54edce9d61aa9f8d09a3cd691706d4f661c912551f83bb
