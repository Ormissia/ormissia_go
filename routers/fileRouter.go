// @File: fileRouter
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 文件的路由配置

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ormissia/go-gin-blog/controller"
	"github.com/ormissia/go-gin-blog/middleware"
)

func FileRouter(e *gin.Engine) {
	fileController := &controller.FileController{}
	//路由组
	private := e.Group("api/file")
	//将登录验证的中间件注册到需要验证的路由
	private.Use(middleware.Auth())
	{
		private.POST("/uploadFile", fileController.UploadImage)
	}
}
