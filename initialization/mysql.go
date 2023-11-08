package initialization

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := Config.Mysql.Dsn()
	if Config.Mysql.DBName == "" {
		DB = nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		DB = nil
	} else {
		mysqldb, _ := db.DB()
		mysqldb.SetMaxIdleConns(Config.Mysql.MaxIdleConns)
		mysqldb.SetMaxOpenConns(Config.Mysql.MaxOpenConns)
		DB = db
	}

}
