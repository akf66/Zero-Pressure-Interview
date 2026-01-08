package initialize

import (
	"net"
	"strconv"
	"zpi/server/cmd/api/config"
	"zpi/server/shared/consts"
	"zpi/server/shared/tools"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.ApiConfigPath)
	if err := v.ReadInConfig(); err != nil {
		hlog.Fatalf("viper read config failed: %v", err)
	}
	if err := v.Unmarshal(&config.GlobalConsulConfig); err != nil {
		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	}

	hlog.Infof("Config Info: %v", config.GlobalConsulConfig)
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port))
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %v", err)
	}
	content, _, err := consulClient.KV().Get(config.GlobalConsulConfig.Key, nil)
	if err != nil {
		hlog.Fatalf("consul kv failed: %s", err.Error())
	}
	err = sonic.Unmarshal(content.Value, &config.GlobalServerConfig)
	if err != nil {
		hlog.Fatalf("server config unmarshal err: %s", err.Error())
	}

	if config.GlobalServerConfig.Host == "" {
		config.GlobalServerConfig.Host, err = tools.GetLocalIPv4Address()
		hlog.Infof("get local ip address: %v", config.GlobalServerConfig.Host)
		if err != nil {
			hlog.Fatalf("get local ipv4 address failed: %s", err.Error())
		}
	}
}
