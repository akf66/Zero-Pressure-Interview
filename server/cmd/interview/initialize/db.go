package initialize

import (
	"fmt"
	"time"
	"zpi/server/cmd/interview/config"
	"zpi/server/shared/consts"
	"zpi/server/shared/dal/sqlentity"
	"zpi/server/shared/dal/sqlfunc"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var Q *sqlfunc.Query

// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf(consts.MySqlDSN, c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		klog.Fatalf("open mysql failed: %s", err.Error())
	}

	// 自动迁移面试相关表
	err = db.AutoMigrate(
		&sqlentity.Interview{},
		&sqlentity.InterviewMessage{},
	)
	if err != nil {
		klog.Fatalf("auto migrate failed: %s", err.Error())
	}

	if err = db.Use(tracing.NewPlugin()); err != nil {
		klog.Fatalf("use tracing plugin failed: %s", err.Error())
	}

	// 初始化 gorm gen Query
	Q = sqlfunc.Use(db)

	return db
}
