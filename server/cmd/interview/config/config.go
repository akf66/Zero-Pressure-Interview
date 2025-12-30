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

// AIConfig AI服务配置
type AIConfig struct {
	APIKey   string `mapstructure:"api_key" json:"api_key"`
	Model    string `mapstructure:"model" json:"model"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
}

// RPCSrvConfig RPC服务配置
type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

// ServerConfig Interview服务配置
type ServerConfig struct {
	Name        string       `mapstructure:"name" json:"name"`
	Host        string       `mapstructure:"host" json:"host"`
	MysqlInfo   MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	RedisInfo   RedisConfig  `mapstructure:"redis" json:"redis"`
	OtelInfo    OtelConfig   `mapstructure:"otel" json:"otel"`
	AIInfo      AIConfig     `mapstructure:"ai" json:"ai"`
	UserSrvInfo RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
}
