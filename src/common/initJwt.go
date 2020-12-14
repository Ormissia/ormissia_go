// @File: initJwt
// @Date: 2020/12/7 8:58
// @Author: 安红豆
// @Description: jwt
package common

import (
	"github.com/Ormissia/ormissia_go/src/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//jwt加密的密钥
var jwtKey = []byte("ormissia_secret")

//token的Claims
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//登录成功后调用此方法发送token
func ReleaseToken(user model.User) (string, error) {
	//token过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	var claims = &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ormissia.com",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

//解析token
func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
