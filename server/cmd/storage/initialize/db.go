package initialize

import (
	"fmt"
	"zpi/server/cmd/storage/config"
	"zpi/server/shared/dal/sqlfunc"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
	Q  *sqlfunc.Query
)

// InitDB 初始化数据库连接
func InitDB() {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		klog.Fatalf("init mysql failed: %s", err.Error())
	}

	// 初始化 gorm gen Query
	Q = sqlfunc.Use(DB)
}
