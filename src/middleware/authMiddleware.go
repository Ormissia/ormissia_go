// @File: authMiddleware
// @Date: 2020/12/12 16:10
// @Author: 安红豆
// @Description: 权限认证
package middleware

import (
	"github.com/Ormissia/ormissia_go/src/model"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
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
		//验证token格式是否合法
		token, claims, err := util.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			util.Band(ctx, "请先登录", nil)
			ctx.Abort()
			return
		}

		//token格式正确
		//验证token是否过期
		if claims.ExpiresAt < time.Now().Unix() {
			util.Band(ctx, "请先登录", nil)
			ctx.Abort()
			return
		}

		//验证userId是否存在
		user, _ := model.SelectUserInfoByUserId(claims.UserId)
		if user.ID == 0 {
			util.Band(ctx, "请先登录", nil)
			ctx.Abort()
			return
		}

		//验证通过将user的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
