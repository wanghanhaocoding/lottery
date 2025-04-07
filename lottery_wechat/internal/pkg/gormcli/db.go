package gormcli

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lottery_wechat/configs"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	db   *gorm.DB
	once sync.Once
)

func openDb() {
	dbConfig := configs.GetGlobalConfig().DbConfig

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.PassWord, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
	log.Infof("connArgs:%s", connArgs)

	var err error
	db, err = gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open db err:%v", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("fetch db err:" + err.Error())
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Duration(dbConfig.MaxIdleTime * int(time.Second)))
}

func GetDb() *gorm.DB {
	once.Do(openDb)
	return db
}
