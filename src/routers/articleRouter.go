// @File: articleRouter
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 文章的路由配置
package routers

import (
	"github.com/Ormissia/ormissia_go/src/controller"
	"github.com/Ormissia/ormissia_go/src/middleware"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(e *gin.Engine) {
	articleController := &controller.ArticleController{}
	//路由组
	private := e.Group("api/article")
	//将登录验证的中间件注册到需要验证的路由
	private.Use(middleware.Auth())
	{
		private.POST("/saveArticle", articleController.SaveArticle)
		private.POST("/deleteArticle", articleController.DeleteArticle)
	}
	public := e.Group("api/article")
	{
		public.POST("/selectArticleById", articleController.SelectArticleById)
		public.POST("/selectArticleByPage", articleController.SelectArticleByPage)
	}
}
