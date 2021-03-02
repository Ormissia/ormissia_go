// @File: userRouter
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 用户相关的路由配置

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ormissia/go-gin-blog/controller"
)

//登录路由
func UserRouter(e *gin.Engine) {
	userController := &controller.UserController{}
	//路由组
	user := e.Group("/api/user")
	{
		//隐藏注册接口Casbin过一阵再写吧
		//user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}
}
