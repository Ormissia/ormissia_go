// @File: blog
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 博客
package model

import "gorm.io/gorm"

//博客属性
type Article struct {
	User User  `gorm:"one2one:user;foreignKey:id"`          //博客所属用户
	Type Type  `gorm:"one2one:type;foreignKey:id"`          //类型
	Tags []Tag `gorm:"many2many:article_tag;foreignKey:id"` //标签
	gorm.Model
	UserId      string `json:"user_id" gorm:"type:varchar(20);not null"` //博客所属用户id
	TypeId      string `json:"type_id" gorm:"type:varchar(20)"`          //类型
	TopImage    string `json:"top_image" gorm:"type:varchar(20)"`        //首图
	Title       string `json:"title" gorm:"type:varchar(20)"`            //标题
	Description string `json:"description" gorm:"type:varchar(20)"`      //描述
	Content     string `json:"content" gorm:"type:longtext"`             //内容
	Visits      int    `json:"visits" gorm:"type:int"`                   //访客数量
	IsDeleted   bool   `json:"is_deleted" gorm:"int"`                    //是否删除（正常、删除）
	IsRecommend bool   `json:"is_recommend" gorm:"int"`                  //是否推荐（推荐、普通）
	IsPublished bool   `json:"is_published" gorm:"int"`                  //是否发布（发布、草稿）
}

//实现gorm的接口，重命名表名
func (Article) TableName() string {
	return "article"
}
