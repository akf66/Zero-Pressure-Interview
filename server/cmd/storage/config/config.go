package config

// MysqlConfig MySQL数据库配置
type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

// ConsulConfig Consul配置
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

// OtelConfig OpenTelemetry配置
type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
	PoolSize int    `mapstructure:"pool_size" json:"pool_size"`
}

// MinIOConfig MinIO对象存储配置
type MinIOConfig struct {
	Endpoint  string `mapstructure:"endpoint" json:"endpoint"`
	AccessKey string `mapstructure:"access_key" json:"access_key"`
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	UseSSL    bool   `mapstructure:"use_ssl" json:"use_ssl"`
	Bucket    string `mapstructure:"bucket" json:"bucket"`
}

// ServerConfig Storage服务配置
type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`
	Host      string      `mapstructure:"host" json:"host"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
	MinIOInfo MinIOConfig `mapstructure:"minio" json:"minio"`
	OtelInfo  OtelConfig  `mapstructure:"otel" json:"otel"`
}
