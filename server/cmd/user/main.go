package main

import (
	"context"
	"net"
	"strconv"
	"zpi/server/cmd/user/config"
	"zpi/server/cmd/user/initialize"
	"zpi/server/cmd/user/pkg/email"
	"zpi/server/cmd/user/pkg/md5"
	"zpi/server/cmd/user/pkg/paseto"
	"zpi/server/cmd/user/pkg/verifycode"
	"zpi/server/shared/consts"
	"zpi/server/shared/dal/sqlfunc"
	"zpi/server/shared/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	// initialization
	initialize.InitLogger()
	initialize.InitConfig()
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitRegistry(Port)
	mdb := initialize.InitDB()
	initialize.InitRedis()
	defer initialize.CloseRedis()

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	tg, err := paseto.NewTokenGenerator(
		config.GlobalServerConfig.PasetoInfo.SecretKey,
		[]byte(config.GlobalServerConfig.PasetoInfo.Implicit))
	if err != nil {
		klog.Fatal(err)
	}

	// 初始化邮件发送器
	emailSender := email.NewEmailSender(
		config.GlobalServerConfig.EmailInfo.SMTPHost,
		config.GlobalServerConfig.EmailInfo.SMTPPort,
		config.GlobalServerConfig.EmailInfo.Username,
		config.GlobalServerConfig.EmailInfo.Password,
		config.GlobalServerConfig.EmailInfo.FromName,
	)

	// 初始化验证码管理器
	vcManager := verifycode.NewVerifyCodeManager(initialize.RedisClient)

	// Create new server.
	srv := userservice.NewServer(&UserServiceImpl{
		EncryptManager:    &md5.EncryptManager{Salt: config.GlobalServerConfig.MysqlInfo.Salt},
		TokenGenerator:    tg,
		UserManager:       &UserManager{Query: sqlfunc.Use(mdb)},
		VerifyCodeManager: vcManager,
		EmailSender:       emailSender,
	},
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err = srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
