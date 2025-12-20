package config

type PasetoConfig struct {
	PubKey   string `mapstructure:"pub_key" json:"pub_key"`
	Implicit string `mapstructure:"implicit" json:"implicit"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

// OtelConfig OpenTelemetry链路追踪配置
type OtelConfig struct {
	EndPoint    string `mapstructure:"endpoint" json:"endpoint"`
	ServiceName string `mapstructure:"service_name" json:"service_name"`
}

// MinIOConfig MinIO对象存储配置
type MinIOConfig struct {
	Endpoint  string `mapstructure:"endpoint" json:"endpoint"`
	AccessKey string `mapstructure:"access_key" json:"access_key"`
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	UseSSL    bool   `mapstructure:"use_ssl" json:"use_ssl"`
	Bucket    string `mapstructure:"bucket" json:"bucket"`
}

// ServerConfig API网关服务配置
type ServerConfig struct {
	Name             string       `mapstructure:"name" json:"name"`
	Host             string       `mapstructure:"host" json:"host"`
	Port             int          `mapstructure:"port" json:"port"`
	ChatToken        string       `mapstructure:"chat_token" json:"chat_token"`
	ProxyURL         string       `mapstructure:"proxy" json:"proxy"`
	PasetoInfo       PasetoConfig `mapstructure:"paseto" json:"paseto"`
	OtelInfo         OtelConfig   `mapstructure:"otel" json:"otel"`
	UserSrvInfo      RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	InterviewSrvInfo RPCSrvConfig `mapstructure:"interview_srv" json:"interview_srv"`
	QuestionSrvInfo  RPCSrvConfig `mapstructure:"question_srv" json:"question_srv"`
	StorageSrvInfo   RPCSrvConfig `mapstructure:"storage_srv" json:"storage_srv"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
