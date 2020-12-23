// @File: articleDao
// @Date: 2020/12/14 14:55
// @Author: 安红豆
// @Description: 文章的数据库层操作
package dao

import (
	"github.com/Ormissia/ormissia_go/src/database"
	"github.com/Ormissia/ormissia_go/src/model"
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
	err = database.DB.Table("article").Where("id = ?", id).First(&article).Error
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
		// TODO 嵌套结构体的查询
		Select("article.id", "article.created_at", "article.updated_at",
			"article.deleted_at", "user_id", "type_id", "top_image", "title",
			"description", "visits", "is_deleted", "is_recommend", "is_published").
		Table("article").
		//连接用户表和类型标签表
		Joins("left join user on article.user_id = user.id").
		Joins("left join type on article.type_id = type.id").
		//分页参数是必传的
		//Limit指定获取记录的最大数量,Offset指定在开始返回记录之前要跳过的记录数量
		Offset((page.PageNum-1)*page.PageSize).Limit(page.PageSize).
		//是否删除、推荐、发布在结构体中会有默认值，因此不需要判断是否为空
		Where("article.is_deleted = ? and article.is_recommend = ? and article.is_published = ?",
			page.IsDeleted, page.IsRecommend, page.IsPublished).
		//根据更新时间排序，如果更新时间为空则以创建时间为准
		Order("ifnull( article.updated_at, article.created_at ) desc")

	//动态拼接查询参数需要判空的动态查询参数
	if page.Title != "" {
		query = query.Where("title like ?", "%"+page.Title+"%")
	}

	//执行查询操作
	err = query.Find(&articles).Error
	return
}

//根据分页参数查询文章总数
func CountArticleByPage(page model.ArticlePage) (count int, err error) {
	//查询文章列表
	query := database.DB.
		//指定查询字段(主要是为了排除content字段)
		// TODO 嵌套结构体的查询
		Select("article.id").Table("article").
		//连接用户表和类型标签表
		Joins("left join user on article.user_id = user.id").
		Joins("left join type on article.type_id = type.id").
		//是否删除、推荐、发布在结构体中会有默认值，因此不需要判断是否为空
		Where("article.is_deleted = ? and article.is_recommend = ? and article.is_published = ?",
			page.IsDeleted, page.IsRecommend, page.IsPublished)

	//动态拼接查询参数需要判空的动态查询参数
	if page.Title != "" {
		query = query.Where("title like ?", "%"+page.Title+"%")
	}
	var articles []model.Article
	//执行查询操作
	err = query.Find(&articles).Error
	count = int(query.RowsAffected)
	return
}

//改
func UpdateArticle(article model.Article) (err error) {
	err = database.DB.Omit("id", "created_at").Updates(&article).Error
	return
}
