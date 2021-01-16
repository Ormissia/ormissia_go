// @File: articleController
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 文章的控制层
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ormissia/go-gin-blog/common"
	"github.com/ormissia/go-gin-blog/dao"
	"github.com/ormissia/go-gin-blog/model"
	"github.com/ormissia/go-gin-blog/util"
)

type ArticleController struct {
}

//TODO 文章标签关系修改BUG
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
	if article.ID == 0 {
		//当id等于0时为新增操作
		//新增时需要插入userId
		//获取Authorization Header
		tokenStr := ctx.GetHeader("token")[7:]
		//验证token格式是否合法
		_, claims, _ := common.ParseToken(tokenStr)
		//验证userId是否存在
		user, _ := dao.SelectUserInfoByUsername(claims.Username)
		article.UserId = user.ID
		err = dao.InsertArticle(article)
	} else {
		//当id不等0时为修改操作
		err = dao.UpdateArticle(article)
	}
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	util.Success(ctx, "保存成功", nil)
}

//删除文章（删除功能为逻辑删除，暂时与修改共用同一个接口）
func (w *ArticleController) DeleteArticle(ctx *gin.Context) {
}

//TODO 查询功能需要添加userId的条件
//根据id查询文章
func (w *ArticleController) SelectArticleById(ctx *gin.Context) {
	//从请求中获取要查询的文章id
	//articleId := ctx.Query("articleId")
	articleId := ctx.PostForm("articleId")
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
	total, err := dao.CountArticleByPage(page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//返回成功
	util.Success(ctx, util.GetCodeMsg(util.HttpSuccess), map[string]interface{}{
		"total":    total,
		"articles": articles,
	})
	return
}
