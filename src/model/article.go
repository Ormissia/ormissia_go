// @File: blog
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 博客
package model

import (
	"gorm.io/gorm"
)

//博客属性
type Article struct {
	User User  `json:"user" gorm:"one2one:user;foreignKey:user_id;references:id"` //博客所属用户
	Type Type  `json:"type" gorm:"one2one:type;foreignKey:type_id;references:id"` //类型
	Tags []Tag `json:"tags" gorm:"many2many:article_tag;foreignKey:id"`           //标签
	gorm.Model
	UserId      string `json:"userId" gorm:"type:varchar(100);not null"` //博客所属用户id
	TypeId      string `json:"typeId" gorm:"type:varchar(100)"`          //类型
	TopImage    string `json:"topImage" gorm:"type:varchar(100)"`        //首图
	Title       string `json:"title" gorm:"type:varchar(100)"`           //标题
	Description string `json:"description" gorm:"type:varchar(100)"`     //描述
	Content     string `json:"content" gorm:"type:longtext"`             //内容
	Visits      int    `json:"visits" gorm:"type:int"`                   //访客数量
	IsDeleted   int    `json:"isDeleted" gorm:"int"`                     //是否删除（0正常、1删除）
	IsRecommend int    `json:"isRecommend" gorm:"int"`                   //是否推荐（0普通、1推荐）
	IsPublished int    `json:"isPublished" gorm:"int"`                   //是否发布（0草稿、1发布）
}

//分页查询的属性
type ArticlePage struct {
	PageNum  int //第几页
	PageSize int //每页数量
	Article
}

//实现gorm的接口，重命名表名
func (Article) TableName() string {
	return "article"
}
