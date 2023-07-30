package DbUtil

import (
	"fmt"
	"github.com/marqstree/gstep/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

// 定义mysql连接
func Setup() {
	var db *gorm.DB

	var err error
	var connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Db.User,
		config.Config.Db.Password,
		config.Config.Db.Host,
		config.Config.Db.Database)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
		//禁用表名复数形式
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         newLogger,
	})

	if err != nil {
		log.Printf("[info] gorm %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("[info] gorm %s", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	Db = db
}

func GetTx() *gorm.DB {
	return Db.Begin()
}
