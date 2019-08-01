package model

import (
	"Shepherd/pkg/config"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// 数据库操作对象
var DB *gorm.DB

// 初始化数据库
func Setup() {
	// 连接数据库
	var err error
	DB, err = gorm.Open(config.Database.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name))
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}
	// 设置表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Database.TablePrefix + defaultTableName
	}
	// 禁用表名多元化
	DB.SingularTable(true)
	// 是否开启debug模式
	if config.Database.Debug {
		// debug 模式
		DB = DB.Debug()
	}
	// 连接池最大连接数
	DB.DB().SetMaxIdleConns(config.Database.MaxIdleConns)
	// 默认打开连接数
	DB.DB().SetMaxOpenConns(config.Database.MaxOpenConns)
	// 开启协程ping MySQL数据库查看连接状态
	go func() {
		for {
			// ping
			err = DB.DB().Ping()
			if err != nil {
				log.Error().Msg(err.Error())
			}
			// 间隔30s ping一次
			time.Sleep(config.Database.PingInterval)
		}
	}()
}

// 关闭连接
func Close() {
	if DB != nil {
		defer DB.Close()
	}
}
