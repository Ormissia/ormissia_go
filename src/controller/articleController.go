// @File: articleController
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 文章的控制层
package controller

import (
	"github.com/Ormissia/ormissia_go/src/dao"
	"github.com/Ormissia/ormissia_go/src/model"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

//保存文章（新增和修改）
func (w *ArticleController) SaveArticle(ctx *gin.Context) {
	article := model.Article{}
	//从请求中获取参数
	err := ctx.Bind(&article)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}

	//执行插入或修改操作
	err = dao.InsertArticle(article)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	util.Success(ctx, "保存成功", nil)
}

//删除文章
func (w *ArticleController) DeleteArticle(ctx *gin.Context) {

}

//根据id查询文章
func (w *ArticleController) SelectArticleById(ctx *gin.Context) {
	//从请求中获取要查询的文章id
	articleId := ctx.Query("id")
	article, err := dao.SelectArticleById(articleId)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	util.Success(ctx, util.GetCodeMsg(util.HttpSuccess), article)
	return
}

//根据分页参数查询文章
func (w *ArticleController) SelectArticleByPage(ctx *gin.Context) {
	page := model.ArticlePage{}
	//从请求中获取参数
	err := ctx.Bind(&page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//根据分页参数查询博客列表
	articles, err := dao.SelectArticleByPage(page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//根据分页参数查询博客总数
	count, err := dao.CountArticleByPage(page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//返回成功
	util.Success(ctx, util.GetCodeMsg(util.HttpSuccess), map[string]interface{}{
		"count":    count,
		"articles": articles,
	})
	return
}
