// @File: userDao
// @Date: 2020/12/14 14:56
// @Author: 安红豆
// @Description:
package dao

import (
	"github.com/ormissia/go-gin-blog/database"
	"github.com/ormissia/go-gin-blog/model"
)

//增
//插入一个新用户
func InsertUser(user model.User) (err error) {
	return database.DB.Create(&user).Error
}

//删

//查
//根据用户id查询用户信息
func SelectUserInfoByUserId(id uint) (user *model.User, err error) {
	//赋值给User
	user = &model.User{}
	//查询用户信息
	err = database.DB.Table("user").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return
}

//根据用户名称查询用户信息
func SelectUserInfoByUsername(username string) (user *model.User, err error) {
	//赋值给User
	user = &model.User{}
	//查询用户信息
	err = database.DB.Table("user").Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return
}

//改
