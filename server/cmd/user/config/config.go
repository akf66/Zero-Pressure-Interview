package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Salt     string `mapstructure:"salt" json:"salt"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type PasetoConfig struct {
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	Implicit  string `mapstructure:"implicit" json:"implicit"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
	PoolSize int    `mapstructure:"pool_size" json:"pool_size"`
}

type EmailConfig struct {
	SMTPHost string `mapstructure:"smtp_host" json:"smtp_host"`
	SMTPPort int    `mapstructure:"smtp_port" json:"smtp_port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	FromName string `mapstructure:"from_name" json:"from_name"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	PasetoInfo PasetoConfig `mapstructure:"paseto" json:"paseto"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	RedisInfo  RedisConfig  `mapstructure:"redis" json:"redis"`
	EmailInfo  EmailConfig  `mapstructure:"email" json:"email"`
	OtelInfo   OtelConfig   `mapstructure:"otel" json:"otel"`
}

type RpcSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
