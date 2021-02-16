// @File: typeRouter
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 类型的路由配置

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ormissia/go-gin-blog/controller"
	"github.com/ormissia/go-gin-blog/middleware"
)

func TypeRouter(e *gin.Engine) {
	typeController := &controller.TypeController{}
	//路由组
	private := e.Group("api/type")
	//将登录验证的中间件注册到需要验证的路由
	private.Use(middleware.Auth())
	{
		private.POST("/saveType", typeController.SaveType)
	}
	public := e.Group("api/type")
	{
		public.POST("/selectTypeByPage", typeController.SelectTypeByPage)
	}
}
