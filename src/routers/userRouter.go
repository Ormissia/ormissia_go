// @File: userRouter
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 用户相关的路由配置
package routers

import (
	"github.com/gin-gonic/gin"
	"ormissia_go/src/controller"
)

//登录路由
func LoginRouter(e *gin.Engine) {
	loginController := &controller.UserController{}
	//登录认证路由组
	loginAuth := e.Group("/api/v1/user")
	{
		loginAuth.POST("/login", loginController.Login)
	}
}
