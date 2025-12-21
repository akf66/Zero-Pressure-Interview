package verifycode

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// 验证码长度
	CodeLength = 6
	// 验证码有效期（5分钟）
	CodeExpiration = 5 * time.Minute
	// 验证码发送限制（1小时内最多5次）
	RateLimitCount  = 5
	RateLimitWindow = 1 * time.Hour
)

// VerifyCodeManager 验证码管理器
type VerifyCodeManager struct {
	redis *redis.Client
}

// NewVerifyCodeManager 创建验证码管理器
func NewVerifyCodeManager(redisClient *redis.Client) *VerifyCodeManager {
	return &VerifyCodeManager{
		redis: redisClient,
	}
}

// GenerateCode 生成6位数字验证码
func (m *VerifyCodeManager) GenerateCode() (string, error) {
	code := ""
	for i := 0; i < CodeLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code += num.String()
	}
	return code, nil
}

// StoreCode 存储验证码到 Redis
// key 格式: verify_code:{type}:{target}:{purpose}
func (m *VerifyCodeManager) StoreCode(ctx context.Context, codeType, target, purpose, code string) error {
	key := fmt.Sprintf("verify_code:%s:%s:%s", codeType, target, purpose)
	return m.redis.Set(ctx, key, code, CodeExpiration).Err()
}

// VerifyCode 验证验证码
func (m *VerifyCodeManager) VerifyCode(ctx context.Context, codeType, target, purpose, code string) (bool, error) {
	key := fmt.Sprintf("verify_code:%s:%s:%s", codeType, target, purpose)

	storedCode, err := m.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil // 验证码不存在或已过期
	}
	if err != nil {
		return false, err
	}

	// 验证成功后删除验证码（一次性使用）
	if storedCode == code {
		_ = m.redis.Del(ctx, key).Err()
		return true, nil
	}

	return false, nil
}

// CheckRateLimit 检查发送频率限制
func (m *VerifyCodeManager) CheckRateLimit(ctx context.Context, target string) (bool, error) {
	key := fmt.Sprintf("rate_limit:verify_code:%s", target)

	count, err := m.redis.Get(ctx, key).Int()
	if err == redis.Nil {
		// 第一次发送
		return true, nil
	}
	if err != nil {
		return false, err
	}

	// 检查是否超过限制
	if count >= RateLimitCount {
		return false, nil
	}

	return true, nil
}

// IncrementRateLimit 增加发送计数
func (m *VerifyCodeManager) IncrementRateLimit(ctx context.Context, target string) error {
	key := fmt.Sprintf("rate_limit:verify_code:%s", target)

	// 增加计数
	count, err := m.redis.Incr(ctx, key).Result()
	if err != nil {
		return err
	}

	// 如果是第一次，设置过期时间
	if count == 1 {
		return m.redis.Expire(ctx, key, RateLimitWindow).Err()
	}

	return nil
}

// GetRemainingTime 获取验证码剩余有效时间
func (m *VerifyCodeManager) GetRemainingTime(ctx context.Context, codeType, target, purpose string) (time.Duration, error) {
	key := fmt.Sprintf("verify_code:%s:%s:%s", codeType, target, purpose)
	return m.redis.TTL(ctx, key).Result()
}

// DeleteCode 删除验证码
func (m *VerifyCodeManager) DeleteCode(ctx context.Context, codeType, target, purpose string) error {
	key := fmt.Sprintf("verify_code:%s:%s:%s", codeType, target, purpose)
	return m.redis.Del(ctx, key).Err()
}
