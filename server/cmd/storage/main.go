package main

import (
	"net"
	"zpi/server/cmd/storage/config"
	"zpi/server/cmd/storage/initialize"
	"zpi/server/shared/consts"
	storage "zpi/server/shared/kitex_gen/storage/storageservice"
	"zpi/server/shared/middleware"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitRedis()
	initialize.InitMinIO()

	// 服务注册
	r, info := initialize.InitRegistry(consts.StorageServicePort)

	// 创建服务实现
	impl := &StorageServiceImpl{
		StorageManager: &StorageManager{
			Query: initialize.Q,
		},
	}

	// 创建服务
	svr := storage.NewServer(
		impl,
		server.WithServiceAddr(
			&net.TCPAddr{
				IP:   net.ParseIP(config.GlobalServerConfig.Host),
				Port: consts.StorageServicePort,
			},
		),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{
			MaxConnections: 2000,
			MaxQPS:         500,
		}),
		server.WithMiddleware(middleware.KitexRecovery()),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.GlobalServerConfig.Name,
		}),
	)

	if err := svr.Run(); err != nil {
		klog.Fatalf("storage service run failed: %s", err.Error())
	}
}
