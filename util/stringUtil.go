// @File: stringUtil
// @Date: 2020/12/14 23:07
// @Author: 安红豆
// @Description: 字符串工具类

package util

import (
	"crypto/md5"
	"fmt"
)

//string加密为MD5
func StringConvertMD5(str string) (md5String string) {
	md5String = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return
}
