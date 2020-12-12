// @File: corsMiddleware
// @Date: 2020/12/6 20:53
// @Author: 安红豆
// @Description: 跨域配置
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//跨域
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//允许的域
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		//缓存请求信息，单位为秒
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		//Header的类型
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//跨域请求是否需要带cookie信息,默认设置为true
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//设置返回格式为JSON
		ctx.Set("content-type", "application/json")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
