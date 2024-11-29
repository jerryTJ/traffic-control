package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(user, passwd, db_url, db_name string) {
	DB, _ = getDbConnection(user, passwd, db_url, db_name)
}
func getDbConnection(user, passwd, db_url, db_name string) (db *gorm.DB, err error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	var mysql_dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, passwd, db_url, db_name)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysql_dsn, // data source name
		DefaultStringSize:         256,       // default size for string fields
		DisableDatetimePrecision:  true,      // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,      // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,      // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,     // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("init db error")
	}
	return db, err
}
