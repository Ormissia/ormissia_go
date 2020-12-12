// @File: recoveryMiddleware
// @Date: 2020/12/11 17:49
// @Author: 安红豆
// @Description: 统一错误处理
package middleware

import (
	"fmt"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				util.Error(ctx, fmt.Sprint(err), nil)
				return
			}
		}()
		ctx.Next()
	}
}
