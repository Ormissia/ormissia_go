// @File: regexpUtil
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 正则工具类
package util

import "regexp"

//判断邮箱地址是否合法
func EmailRegexp(email string) bool {
	matched, _ := regexp.MatchString(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`, email)
	return matched
}
