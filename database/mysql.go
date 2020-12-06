// @File: mysql
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 初始化数据库连接

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

//初始化连接
func InitMySql() error {
	dsn := "root:pass@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	} else {
		//创建数据库连接成功，并将连接返回
		DB = db
		return nil
	}
}
