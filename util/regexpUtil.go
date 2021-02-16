// @File: regexpUtil
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 正则工具类

package util

import "regexp"

//用来匹配的正则表达式常量
const (
	//IP(1.0.0.0~255.255.255.255)
	RegexpExpressionIP = `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	//域名
	RegexpExpressionDomainName = `[a-zA-Z0-9][-a-zA-Z0-9]{1,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{1,62})+`
	//邮箱
	RegexpExpressionEmail = `([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`
)

//正则匹配，第一个参数为用来匹配的正则表达式，第二个参数为需要匹配的字符串
func StringRegexp(regexpString, toBeTestString string) (matched bool) {
	matched, _ = regexp.MatchString(regexpString, toBeTestString)
	return
}

//使用闭包创建正则匹配方法
func RegexpBase(regexpString string) func(toBeTestString string) (matched bool) {
	return func(toBeTestString string) (matched bool) {
		matched, _ = regexp.MatchString(regexpString, toBeTestString)
		return
	}
}
