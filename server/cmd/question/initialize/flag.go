package initialize

import (
	"flag"
	"zpi/server/shared/consts"
	"zpi/server/shared/tools"

	"github.com/cloudwego/kitex/pkg/klog"
)

// InitFlag to init flag
func InitFlag() int {
	Port := flag.Int(consts.PortFlagName, 0, consts.PortFlagUsage)
	// Parsing flags and if Port is 0 , then will automatically get an empty Port.
	flag.Parse()
	if *Port == 0 {
		*Port, _ = tools.GetFreePort()
	}
	klog.Info("port: ", *Port)
	return *Port
}
