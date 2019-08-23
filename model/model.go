package model

import (
	"Shepherd/pkg/config"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(config.Database.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name))
	if err != nil {
		log.Fatal().Msg(err.Error())
		panic(err.Error())
	}

	defer db.Close()

	// 全局禁用表名复数
	db.SingularTable(true)

	// 是否开启debug模式
	if config.Database.Debug {
		db = db.Debug()
	}

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 连接池最大连接数
	db.DB().SetMaxIdleConns(config.Database.MaxIdleConns)
	// 默认打开连接数
	db.DB().SetMaxOpenConns(config.Database.MaxOpenConns)

	// 开启协程ping MySQL数据库查看连接状态
	go func() {
		for {
			err = db.DB().Ping()
			if err != nil {
				log.Error().Msg(err.Error())
			}
			// 间隔Xs ping一次
			time.Sleep(config.Database.PingInterval)
		}
	}()
}
