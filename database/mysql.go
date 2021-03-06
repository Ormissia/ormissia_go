// @File: mysql
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 初始化数据库连接

package database

import (
	"github.com/ormissia/go-gin-blog/model"
	"github.com/ormissia/go-gin-blog/util"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

//初始化连接
func InitMySql() *gorm.DB {
	host := viper.GetString(util.ConfigDataSourceHost)
	port := viper.GetString(util.ConfigDataSourcePort)
	database := viper.GetString(util.ConfigDataSourceDatabase)
	username := viper.GetString(util.ConfigDataSourceUsername)
	password := viper.GetString(util.ConfigDataSourcePassword)
	charset := viper.GetString(util.ConfigDataSourceCharset)
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//配置打印日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	//自动迁移
	err = db.AutoMigrate(
		&model.User{},
		&model.Type{},
		&model.Tag{},
		&model.Article{},
	)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	sqlDB, err := db.DB()
	//设置连接池中最大闲置连接数
	sqlDB.SetMaxIdleConns(10)

	//设置数据库的最大连接数量
	sqlDB.SetMaxOpenConns(100)

	//设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(time.Second * 10)
	//创建数据库连接成功，并将连接返回
	DB = db
	return db
}
