package config

import (
	"zpi/server/shared/kitex_gen/interview/interviewservice"
	"zpi/server/shared/kitex_gen/question/questionservice"
	"zpi/server/shared/kitex_gen/storage/storageservice"
	"zpi/server/shared/kitex_gen/user/userservice"
)

var (
	// GlobalServerConfig API网关服务配置
	GlobalServerConfig ServerConfig

	// GlobalConsulConfig 注册中心的配置
	GlobalConsulConfig ConsulConfig

	// RPC客户端实例
	GlobalUserClient      userservice.Client      // 用户服务客户端
	GlobalInterviewClient interviewservice.Client // 面试服务客户端
	GlobalQuestionClient  questionservice.Client  // 题库服务客户端
	GlobalStorageClient   storageservice.Client   // 存储服务客户端
)
