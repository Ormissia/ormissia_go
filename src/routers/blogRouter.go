// @File: blogRouter
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description:
package routers

import (
	"github.com/Ormissia/ormissia_go/src/controller"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(e *gin.Engine) {
	articleController := &controller.ArticleController{}
	//文章路由组
	article := e.Group("api/article")
	{
		article.POST("/saveArticle", articleController.SaveArticle)
		article.POST("/deleteArticle", articleController.DeleteArticle)
		article.POST("/selectBlogById", articleController.SelectBlogById)
		article.POST("/selectBlogByPage", articleController.SelectBlogByPage)
	}
}
