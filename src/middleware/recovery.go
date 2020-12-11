// @File: recovery
// @Date: 2020/12/11 17:49
// @Author: 安红豆
// @Description: 统一错误处理
package middleware

import (
	"fmt"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				util.Error(c, fmt.Sprint(err), nil)
				return
			}
		}()
		c.Next()
	}
}
