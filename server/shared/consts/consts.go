package consts

const (
	ZPI = "ZPI"

	SystemAdmin = "SystemAdmin" // 系统管理员
	BackAdmin   = "BackAdmin"   //运营管理员
	User        = "User"        //普通用户

	AccountID = "accountID"

	ApiConfigPath       = "./server/cmd/api/config.yaml"
	UserConfigPath      = "./server/cmd/user/config.yaml"
	InterviewConfigPath = "./server/cmd/interview/config.yaml"
	StorageConfigPath   = "./server/cmd/storage/config.yaml"

	CorsAddress = "http://localhost:3000"

	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	// logpath
	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"

	TCP = "tcp"

	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"

	// 服务端口
	UserServicePort      = 8881
	InterviewServicePort = 8882
	QuestionServicePort  = 8883
	StorageServicePort   = 8884
)
