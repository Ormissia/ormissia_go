// @File: cors
// @Date: 2020/12/6 20:53
// @Author: 安红豆
// @Description: 跨域配置

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//跨域
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		//请求方法
		method := context.Request.Method
		//请求头
		origin := context.Request.Header.Get("")
		//生命请求头keys
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			//允许所有域
			context.Header("Access-Control-Allow-Origin", "*")
			//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//Header的类型
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//允许跨域设置
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			//缓存请求信息，单位为秒
			context.Header("Access-Control-Max-Age", "172800")
			//跨域请求是否需要带cookie信息,默认设置为true
			context.Header("Access-Control-Allow-Credentials", "true")
			//设置返回格式为JSON
			context.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}

		//处理请求，调用后续的处理函数
		context.Next()
	}
}
