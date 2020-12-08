// @File: mysql
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 初始化数据库连接
package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ormissia_go/src/util"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	//创建数据库连接成功，并将连接返回
	DB = db
	return db
}

//获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
