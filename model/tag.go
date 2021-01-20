// @File: tag
// @Date: 2020/12/10 16:40
// @Author: 安红豆
// @Description: 标签
package model

import "gorm.io/gorm"

//标签属性
type Tag struct {
	//标签下的文章
	Articles []Article `json:"articles" gorm:"many2many:article_tag;foreignKey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	gorm.Model
	TagName   string `json:"tagName" gorm:"type:varchar(100);not null"` //标签名字
	IsDeleted int    `json:"isDeleted" gorm:"int"`                      //是否删除（1删除、2正常）
}

//分页查询的属性
type TagPage struct {
	PageNum  int //第几页
	PageSize int //每页数量
	Tag
}

//实现gorm的接口，重命名表名
func (Tag) TableName() string {
	return "tag"
}
