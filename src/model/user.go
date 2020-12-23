// @File: user
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 用户
package model

import (
	"github.com/Ormissia/ormissia_go/src/util"
	"gorm.io/gorm"
)

//用户属性
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100);not null"` //用户名
	Password string `json:"password" gorm:"type:varchar(100);not null"` //密码
	Email    string `json:"email" gorm:"type:varchar(100)"`             //邮箱
	Phone    string `json:"phone" gorm:"type:varchar(100)"`             //手机号
	Avatar   string `json:"avatar" gorm:"type:varchar(100)"`            //头像
	Role     int    `json:"role" gorm:"type:int"`                       //角色
}

//分页查询的属性
type UserPage struct {
	PageNum  int //第几页
	PageSize int //每页数量
	User
}

//实现gorm的接口，重命名表名
func (User) TableName() string {
	return "user"
}

//新增用户的钩子函数，密码加密
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = util.StringConvertMD5(u.Password)
	return nil
}
