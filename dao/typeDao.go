// @File: typeDao
// @Date: 2020/12/23 17:55
// @Author: 安红豆
// @Description: 类型的数据库层操作
package dao

import (
	"github.com/ormissia/go-gin-blog/database"
	"github.com/ormissia/go-gin-blog/model"
)

//增
func InsertType(articleType model.Type) (err error) {
	err = database.DB.Create(&articleType).Error
	return
}

//删

//查
//根据id查询类型详情
func SelectTypeById(id string) (articleType *model.Type, err error) {
	//赋值给articleType
	articleType = &model.Type{}
	//查询类型信息
	err = database.DB.Table("type").Where("id = ?", id).First(&articleType).Error
	if err != nil {
		return nil, err
	}
	return
}

//根据分页参数查询类型列表
func SelectTypeByPage(page model.TypePage) (articleTypes []model.Type, err error) {
	//查询类型列表
	query := database.DB.
		Select("type.id", "type.created_at", "type.updated_at",
			"type.deleted_at", "type_name", "count(article.id) as articles").
		Table("type").
		Preload("Articles").
		Joins("left join article on type.id = article.type_id").
		Group("id").
		//分页参数是必传的
		//Limit指定获取记录的最大数量,Offset指定在开始返回记录之前要跳过的记录数量
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		//根据类型下文章数量排序，文章数量相同时按照名称排序
		Order("articles desc, type_name desc")

	//动态拼接查询参数需要判空的动态查询参数
	if page.IsDeleted != 0 {
		query = query.Where("type.is_deleted = ?", page.IsDeleted)
	}
	if page.TypeName != "" {
		query = query.Where("type_name like ?", "%"+page.TypeName+"%")
	}

	//执行查询操作
	err = query.Find(&articleTypes).Error
	return
}

//根据分页参数查询文章总数
func CountTypeByPage(page model.TypePage) (total int, err error) {
	//查询文章列表
	query := database.DB.
		Select("type.id").Table("type")

	//动态拼接查询参数需要判空的动态查询参数
	if page.IsDeleted != 0 {
		query = query.Where("is_deleted = ?", page.IsDeleted)
	}
	if page.TypeName != "" {
		query = query.Where("type_name like ?", "%"+page.TypeName+"%")
	}
	var articleTypes []model.Type
	//执行查询操作
	err = query.Find(&articleTypes).Error
	total = int(query.RowsAffected)
	return
}

//改
func UpdateType(articleTypes model.Type) (err error) {
	err = database.DB.Omit("id", "created_at").Updates(&articleTypes).Error
	return
}
