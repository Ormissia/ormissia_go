// @File: constantHttp
// @Date: 2020/12/14 8:32
// @Author: 安红豆
// @Description: 后台错误常量

package util

//状态码常量
const (
	HttpSuccess = 200
	HttpError   = 500

	//code=1000...用户模块的错误
	ErrorUsernameUsed  = 1001 //用户名已存在
	ErrorPasswordWrong = 1002 //密码错误
	ErrorUserNotExist  = 1003 //用户不存在
	ErrorIllegalInput  = 1004 //输入不合法
	ErrorTokenWrong    = 1011 //token错误
	ErrorTokenOverTime = 1012 //token过期

	SuccessUserRegister = 1050 //注册成功
	//code=2000...文章模块的错误
	ErrorArticleNotFound = 2001 //文章未找到

	//code=3000...分类模块的错误

	//code=4000...标签模块的错误
)

var codeMsg = map[int]string{
	HttpSuccess:        "OK",
	HttpError:          "Error",
	ErrorUsernameUsed:  "用户名已存在",
	ErrorPasswordWrong: "密码错误",
	ErrorUserNotExist:  "用户不存在",
	ErrorIllegalInput:  "输入不合法",
	ErrorTokenWrong:    "token错误",
	ErrorTokenOverTime: "token过期",

	SuccessUserRegister: "注册成功",

	ErrorArticleNotFound: "文章未找到",
}

func GetCodeMsg(code int) string {
	return codeMsg[code]
}
