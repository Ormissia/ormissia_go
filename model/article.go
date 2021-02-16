// @File: blog
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 文章

package model

import (
	"gorm.io/gorm"
)

//文章属性
type Article struct {
	//博客所属用户
	User User `json:"user" gorm:"one2one:user;foreignKey:user_id;references:id"`
	//类型
	Type Type `json:"type" gorm:"one2one:type;foreignKey:type_id;references:id"`
	//标签
	Tags []Tag `json:"tags" gorm:"many2many:article_tag;foreignKey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	gorm.Model
	UserId      uint   `json:"userId" gorm:"type:int;not null"`      //博客所属用户id
	TypeId      uint   `json:"typeId" gorm:"type:int"`               //类型
	TopImage    string `json:"topImage" gorm:"type:varchar(255)"`    //首图
	Title       string `json:"title" gorm:"type:varchar(100)"`       //标题
	Description string `json:"description" gorm:"type:varchar(255)"` //描述
	Content     string `json:"content" gorm:"type:longtext"`         //内容
	Visits      int    `json:"visits" gorm:"type:int"`               //访客数量
	//因为Golang中bool类型默认为false，当一个实体类中有多个bool类型的值需要更新时，无法判断前端没有传值还是是前端传的false，并且每次修改都要求前端传所有bool类型的属性显然也不合理
	//所以采用int类型的变量来表示这些属性的不同状态
	//而在Golang中int类型的默认值为0，因此属性为0值时也无法判断是前端没有传值还是传的0值
	//故，这些属性的状态标识应该从1开始，0表示为空（即前端未传值），1表示true，2表示false
	IsDeleted   int `json:"isDeleted" gorm:"int"`   //是否删除（1删除、2正常）
	IsRecommend int `json:"isRecommend" gorm:"int"` //是否推荐（1推荐、2普通）
	IsPublished int `json:"isPublished" gorm:"int"` //是否发布（1发布、2草稿）
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
