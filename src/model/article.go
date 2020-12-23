// @File: blog
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 博客
package model

import (
	"gorm.io/gorm"
	"time"
)

//博客属性
type Article struct {
	User User  `gorm:"one2one:user;foreignKey:user_id;references:id"` //博客所属用户
	Type Type  `gorm:"one2one:type;foreignKey:type_id;references:id"` //类型
	Tags []Tag `gorm:"many2many:article_tag;foreignKey:id"`           //标签
	gorm.Model
	UserId      string `json:"user_id" gorm:"type:varchar(100);not null"` //博客所属用户id
	TypeId      string `json:"type_id" gorm:"type:varchar(100)"`          //类型
	TopImage    string `json:"top_image" gorm:"type:varchar(100)"`        //首图
	Title       string `json:"title" gorm:"type:varchar(100)"`            //标题
	Description string `json:"description" gorm:"type:varchar(100)"`      //描述
	Content     string `json:"content" gorm:"type:longtext"`              //内容
	Visits      int    `json:"visits" gorm:"type:int"`                    //访客数量
	IsDeleted   bool   `json:"is_deleted" gorm:"int;default:false"`       //是否删除（正常、删除）
	IsRecommend bool   `json:"is_recommend" gorm:"int;default:false"`     //是否推荐（推荐、普通）
	IsPublished bool   `json:"is_published" gorm:"int;default:false"`     //是否发布（发布、草稿）
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

// TODO 更新时候创建时间为零值会有问题
//保存之前设置创建和更新时间
func (a Article) BeforeCreate(tx *gorm.DB) (err error) {
	//判断新增还是修改，并赋默认值
	if a.ID == 0 {
		//id为0的时候是新增
		a.CreatedAt = time.Now()
	} else {
		//id不为0的时候是修改
		a.UpdatedAt = time.Now()
	}
	return
}
