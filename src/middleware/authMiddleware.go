// @File: authMiddleware
// @Date: 2020/12/12 16:10
// @Author: 安红豆
// @Description: 权限认证
package middleware

import (
	"github.com/Ormissia/ormissia_go/src/dao"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//检查token
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Authorization Header
		tokenStr := ctx.GetHeader("Authorization")

		//判断是否有token
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer") {
			util.Band(ctx, "请先登录", nil)
			ctx.Abort()
			return
		}

		tokenStr = tokenStr[7:]
		//判断token格式是否合法
		token, claims, err := util.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//token格式正确，获取userId
		userId := claims.UserId
		user, _ := dao.SelectUserInfoByUserId(userId)
		if user.UserId == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 用户存在 将user 的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
