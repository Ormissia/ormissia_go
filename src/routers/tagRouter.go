// @File: tagRouter
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 标签的路由配置
package routers

import (
	"github.com/Ormissia/ormissia_go/src/controller"
	"github.com/Ormissia/ormissia_go/src/middleware"
	"github.com/gin-gonic/gin"
)

func TagRouter(e *gin.Engine) {
	tagController := &controller.TagController{}
	//路由组
	private := e.Group("api/tag")
	//将登录验证的中间件注册到需要验证的路由
	private.Use(middleware.Auth())
	{
		private.POST("/addTag", tagController.SaveTag)
	}
	public := e.Group("api/tag")
	{
		public.POST("/selectTagByPage", tagController.SelectTagByPage)
	}
}
