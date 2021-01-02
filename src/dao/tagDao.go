// @File: tagDao
// @Date: 2020/12/23 17:55
// @Author: 安红豆
// @Description: 标签的数据库层操作
package dao

import (
	"github.com/Ormissia/ormissia_go/src/database"
	"github.com/Ormissia/ormissia_go/src/model"
)

//增
func InsertTag(articleTag model.Tag) (err error) {
	err = database.DB.Save(&articleTag).Error
	return
}

//删

//查
//根据id查询标签详情
func SelectTagById(id string) (articleTag *model.Tag, err error) {
	//赋值给articleTag
	articleTag = &model.Tag{}
	//查询标签信息
	err = database.DB.Table("tag").Where("id = ?", id).First(&articleTag).Error
	if err != nil {
		return nil, err
	}
	return
}

//根据分页参数查询标签列表
func SelectTagByPage(page model.TagPage) (articleTags []model.Tag, err error) {
	//查询标签列表
	query := database.DB.
		Select("id", "created_at", "updated_at",
			"deleted_at", "tag_name", "count(article_id) as articles").
		Table("tag").
		Preload("Articles").
		Joins("left join article_tag on tag.id = article_tag.tag_id").
		Group("id").
		//分页参数是必传的
		//Limit指定获取记录的最大数量,Offset指定在开始返回记录之前要跳过的记录数量
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		//根据标签下文章数量排序，文章数量相同时按照名称排序
		Order("articles desc, tag_name desc")

	//动态拼接查询参数需要判空的动态查询参数
	if page.IsDeleted != 0 {
		query = query.Where("is_deleted = ?", page.IsDeleted)
	}
	if page.TagName != "" {
		query = query.Where("name like ?", "%"+page.TagName+"%")
	}

	//执行查询操作
	err = query.Find(&articleTags).Error
	return
}

//根据分页参数查询文章总数
func CountTagByPage(page model.TagPage) (total int, err error) {
	//查询文章列表
	query := database.DB.
		Select("tag.id").Table("tag")

	//动态拼接查询参数需要判空的动态查询参数
	if page.IsDeleted != 0 {
		query = query.Where("is_deleted = ?", page.IsDeleted)
	}
	if page.TagName != "" {
		query = query.Where("title like ?", "%"+page.TagName+"%")
	}
	var articleTags []model.Tag
	//执行查询操作
	err = query.Find(&articleTags).Error
	total = int(query.RowsAffected)
	return
}

//改
func UpdateTag(articleTags model.Tag) (err error) {
	err = database.DB.Omit("id", "created_at").Updates(&articleTags).Error
	return
}
