package initialize

import (
	"encoding/json"
	"fmt"
	"zpi/server/cmd/interview/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./server/cmd/interview/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalConsulConfig); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	klog.Infof("Config Info: %v", config.GlobalConsulConfig)

	// 从 Consul 读取配置
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		config.GlobalConsulConfig.Host,
		config.GlobalConsulConfig.Port)

	consulClient, err := api.NewClient(cfg)
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}

	content, _, err := consulClient.KV().Get(config.GlobalConsulConfig.Key, nil)
	if err != nil {
		klog.Fatalf("consul kv failed: %s", err.Error())
	}

	err = json.Unmarshal(content.Value, &config.GlobalServerConfig)
	if err != nil {
		klog.Fatalf("consul kv unmarshal failed: %s", err.Error())
	}
	klog.Infof("Server Config: %v", config.GlobalServerConfig)
}
