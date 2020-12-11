// @File: user
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 用户
package models

//用户属性
type User struct {
	UserId       string //用户id
	UserRoleId   string //用户权限id
	Username     string //用户名
	Password     string //密码
	Email        string //邮箱
	PhoneNumber  string //手机号
	HeadPortrait string //头像
}
