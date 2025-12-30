package initialize

import (
	"os"
	"path"
	"time"
	"zpi/server/shared/consts"

	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger 初始化日志
func InitLogger() {
	// 创建日志目录
	logFilePath := consts.KlogFilePath
	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		panic(err)
	}

	// 配置日志文件
	logFileName := path.Join(logFilePath, "question-"+time.Now().Format("2006-01-02")+".log")
	writer := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    10,   // 最大10MB
		MaxBackups: 5,    // 最多保留5个备份
		MaxAge:     30,   // 最多保留30天
		Compress:   true, // 压缩
	}

	logger := kitexlogrus.NewLogger()
	logger.SetOutput(writer)
	logger.SetLevel(klog.LevelDebug)

	klog.SetLogger(logger)
}
