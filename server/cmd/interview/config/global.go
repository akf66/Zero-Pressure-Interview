package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	// GlobalServerConfig Interview服务配置
	GlobalServerConfig ServerConfig

	// GlobalConsulConfig Consul配置
	GlobalConsulConfig ConsulConfig

	// DB 数据库连接
	DB *gorm.DB

	// RedisClient Redis客户端
	RedisClient *redis.Client
)
