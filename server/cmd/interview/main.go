package main

import (
	"flag"
	"zpi/server/cmd/interview/config"
	"zpi/server/cmd/interview/initialize"
	"zpi/server/shared/consts"
	"zpi/server/shared/kitex_gen/interview/interviewservice"
	"zpi/server/shared/middleware"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	config.DB = initialize.InitDB()
	config.RedisClient = initialize.InitRedis()

	// 解析命令行参数
	var port int
	flag.IntVar(&port, consts.PortFlagName, 8503, consts.PortFlagUsage)
	flag.Parse()

	// 初始化服务注册
	r, info := initialize.InitRegistry(port)

	// 初始化 OpenTelemetry
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)

	// 创建服务实现
	impl := &InterviewServiceImpl{
		InterviewManager: &InterviewManager{
			Query: initialize.Q,
		},
	}

	// 创建服务
	svr := interviewservice.NewServer(
		impl,
		server.WithServiceAddr(info.Addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithMiddleware(middleware.KitexRecovery()),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	klog.Infof("Interview service starting on port %d...", port)
	if err := svr.Run(); err != nil {
		klog.Fatalf("Interview service run failed: %v", err)
	}
}
