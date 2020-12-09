// @File: initJwt
// @Date: 2020/12/7 8:58
// @Author: 安红豆
// @Description: jwt
package util

import (
	"github.com/dgrijalva/jwt-go"
	"ormissia_go/src/models"
)

//jwt加密的密钥
var jwtKwy = []byte("ormissia_secret")

//token的Claims
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//登录成功后调用此方法发送token
func ReleaseToken(user models.User) (string, error) {
	return "", nil
}
