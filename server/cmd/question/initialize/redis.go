package initialize

import (
	"context"
	"fmt"
	"zpi/server/cmd/question/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

// InitRedis 初始化Redis连接
func InitRedis() *redis.Client {
	c := config.GlobalServerConfig.RedisInfo
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.DB,
		PoolSize: c.PoolSize,
	})

	// 测试连接
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		klog.Fatalf("redis connect failed: %s", err.Error())
	}

	klog.Info("Redis connected successfully")
	return client
}
