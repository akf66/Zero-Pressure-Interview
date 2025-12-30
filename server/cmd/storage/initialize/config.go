package initialize

import (
	"encoding/json"
	"net"
	"strconv"
	"zpi/server/cmd/storage/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalConsulConfig); err != nil {
		klog.Fatalf("unmarshal config failed: %s", err.Error())
	}

	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port),
	)

	consulClient, err := api.NewClient(cfg)
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}

	content, _, err := consulClient.KV().Get(config.GlobalConsulConfig.Key, nil)
	if err != nil {
		klog.Fatalf("consul kv get failed: %s", err.Error())
	}

	if err := json.Unmarshal(content.Value, &config.GlobalServerConfig); err != nil {
		klog.Fatalf("unmarshal config failed: %s", err.Error())
	}
}
