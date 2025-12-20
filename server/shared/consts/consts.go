package consts

const (
	ZPI = "ZPI"

	Admin         = "Admin"
	User          = "User"
	AccountID     = "accountID"
	ApiConfigPath = "./server/cmd/api/config.yaml"

	CorsAddress = "http://localhost:3000"

	// logpath
	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"

	TCP = "tcp"
)
