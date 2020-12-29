// @File: response
// @Date: 2020/12/11 11:29
// @Author: 安红豆
// @Description: 统一返回结果
package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//统一封装返回结果
func Response(ctx *gin.Context, httpState int, code int, msg string, data interface{}, token string) {
	ctx.JSON(httpState, gin.H{
		"code":  code,
		"msg":   msg,
		"data":  data,
		"token": token,
	})
}

//返回成功
func Success(ctx *gin.Context, msg string, data interface{}) {
	Response(ctx, http.StatusOK, HttpSuccess, msg, data, "")
}

//返回错误
func Error(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusInternalServerError, HttpError, msg, data, "")
}

//需要登录
func Band(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusUnauthorized, http.StatusUnauthorized, msg, data, "")
}

//参数错误
func ParameterError(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, msg, data, "")
}
