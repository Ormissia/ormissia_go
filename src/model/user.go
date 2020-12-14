// @File: user
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 用户
package model

import (
	"github.com/Ormissia/ormissia_go/src/database"
	"gorm.io/gorm"
)

//用户属性
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(20);not null"` //用户名
	Password string `json:"password" gorm:"type:varchar(20);not null"` //密码
	Email    string `json:"email" gorm:"type:varchar(20)"`             //邮箱
	Phone    string `json:"phone" gorm:"type:varchar(20)"`             //手机号
	Avatar   string `json:"avatar" gorm:"type:varchar(20)"`            //头像
	Role     int    `json:"role" gorm:"type:int"`                      //头像
}

//实现gorm的接口，重命名表名
func (User) TableName() string {
	return "user"
}

//增
//插入一个新用户
func InsertUser(user User) (err error) {
	return database.DB.Create(user).Error
}

//删

//查
//根据用户id查询用户信息
func SelectUserInfoByUserId(id uint) (user *User, err error) {
	//赋值给User
	user = &User{}
	//查询用户信息
	err = database.DB.Table("user").Where("id = ?", id).Find(user).Error
	if err != nil {
		return nil, err
	}
	return
}

//根据用户名称查询用户信息
func SelectUserInfoByUsername(username string) (user *User, err error) {
	//赋值给User
	user = &User{}
	//查询用户信息
	err = database.DB.Table("user").Where("username = ?", username).Find(user).Error
	if err != nil {
		return nil, err
	}
	return
}

//改
