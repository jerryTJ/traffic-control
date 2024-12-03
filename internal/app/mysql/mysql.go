package mysql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	UserName     = "root"
	PassWord     = "Admin@123"
	Host         = "host.docker.internal"
	Port         = 3306
	Database     = "traffic"
	MaxLifetime  = 60 * time.Second
	MaxIdletime  = 30 * time.Second
	MaxOpenconns = 6
	MaxIdleconns = 2
	Dialect      = "mysql"
)

func CreateDB() *gorm.DB {
	var db *gorm.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Asia%%2FShanghai",
		UserName, PassWord, Host, Port, Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info), // 设置日志级别
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	// 获取底层的 *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database handle:", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(MaxOpenconns)   // 设置最大连接数
	sqlDB.SetMaxIdleConns(MaxIdleconns)   // 设置空闲连接数
	sqlDB.SetConnMaxLifetime(MaxLifetime) // 设置连接最大存活时间
	sqlDB.SetConnMaxIdleTime(MaxIdletime) // 设置连接最大空闲时间
	stats := sqlDB.Stats()
	log.Printf("Open connections: %d, In use: %d, Idle: %d\n", stats.OpenConnections, stats.InUse, stats.Idle)
	return db
}
