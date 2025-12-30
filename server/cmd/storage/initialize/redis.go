package initialize

import (
	"context"
	"fmt"
	"zpi/server/cmd/storage/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitRedis 初始化Redis连接
func InitRedis() {
	c := config.GlobalServerConfig.RedisInfo
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.DB,
		PoolSize: c.PoolSize,
	})

	if _, err := RedisClient.Ping(context.Background()).Result(); err != nil {
		klog.Fatalf("init redis failed: %s", err.Error())
	}
}
