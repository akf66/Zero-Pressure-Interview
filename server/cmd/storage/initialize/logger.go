package initialize

import (
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

// InitLogger 初始化日志
func InitLogger() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	klog.SetOutput(os.Stdout)
}
