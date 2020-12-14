// @File: user
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 用户
package model

import (
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
