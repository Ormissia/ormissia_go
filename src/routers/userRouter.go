// @File: userRouter
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 用户相关的路由配置
package routers

import (
	"github.com/Ormissia/ormissia_go/src/controller"
	"github.com/gin-gonic/gin"
)

//登录路由
func UserRouter(e *gin.Engine) {
	userController := &controller.UserController{}
	//用户路由组
	user := e.Group("/api/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}
}
