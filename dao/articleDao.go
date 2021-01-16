// @File: articleDao
// @Date: 2020/12/14 14:55
// @Author: 安红豆
// @Description: 文章的数据库层操作
package dao

import (
	"github.com/ormissia/go-gin-blog/database"
	"github.com/ormissia/go-gin-blog/model"
)

//增
func InsertArticle(article model.Article) (err error) {
	err = database.DB.Save(&article).Error
	return
}

//删

//查
//根据id查询文章详情
func SelectArticleById(id string) (article *model.Article, err error) {
	//赋值给article
	article = &model.Article{}
	//查询文章信息
	err = database.DB.Table("article").
		Preload("User").
		Preload("Type").
		Preload("Tags").
		Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}
	return
}

//根据分页参数查询文章列表
func SelectArticleByPage(page model.ArticlePage) (articles []model.Article, err error) {
	//查询文章列表
	query := database.DB.
		//指定查询字段(主要是为了排除content字段)
		Select("article.id", "article.created_at", "article.updated_at",
			"article.deleted_at", "user_id", "type_id", "top_image", "title",
			"description", "visits", "is_deleted", "is_recommend", "is_published").
		Table("article").
		//连接用户表和类型标签表
		Preload("User").
		Preload("Type").
		Preload("Tags").
		//分页参数是必传的
		//Limit指定获取记录的最大数量,Offset指定在开始返回记录之前要跳过的记录数量
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		//根据创建时间排序
		Order("article.created_at desc")

	//动态拼接查询参数需要判空的动态查询参数
	if page.IsDeleted != 0 {
		query = query.Where("article.is_deleted = ?", page.IsDeleted)
	}
	if page.IsRecommend != 0 {
		query = query.Where("article.is_recommend = ?", page.IsRecommend)
	}
	if page.IsPublished != 0 {
		query = query.Where("article.is_published = ?", page.IsPublished)
	}
	if page.Title != "" {
		query = query.Where("title like ?", "%"+page.Title+"%")
	}

	//执行查询操作
	err = query.Find(&articles).Error
	return
}

//根据分页参数查询文章总数
func CountArticleByPage(page model.ArticlePage) (total int, err error) {
	//查询文章列表
	query := database.DB.
		//指定查询字段(主要是为了排除content字段)
		Select("article.id").Table("article")

	//动态拼接查询参数需要判空的动态查询参数
	if page.IsDeleted != 0 {
		query = query.Where("article.is_deleted = ?", page.IsDeleted)
	}
	if page.IsRecommend != 0 {
		query = query.Where("article.is_recommend = ?", page.IsRecommend)
	}
	if page.IsPublished != 0 {
		query = query.Where("article.is_published = ?", page.IsPublished)
	}
	if page.Title != "" {
		query = query.Where("title like ?", "%"+page.Title+"%")
	}
	var articles []model.Article
	//执行查询操作
	err = query.Find(&articles).Error
	total = int(query.RowsAffected)
	return
}

//改
func UpdateArticle(article model.Article) (err error) {
	err = database.DB.Omit("id", "created_at").Updates(&article).Error
	return
}
