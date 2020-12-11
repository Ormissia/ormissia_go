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
func Response(ctx *gin.Context, httpState int, code int, data gin.H, msg string) {
	ctx.JSON(httpState, gin.H{"code:": code, "data": data, "msg": msg})
}

//返回成功
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

//返回失败
func Fail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}

//返回错误
func Error(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusInternalServerError, 500, data, msg)
}
