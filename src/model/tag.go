// @File: tag
// @Date: 2020/12/10 16:40
// @Author: 安红豆
// @Description: 标签
package model

import "gorm.io/gorm"

//标签属性
type Tag struct {
	Article []Article `json:"article" gorm:"-"` //标签下的文章
	gorm.Model
	TagName string `json:"tag_name" gorm:"type:varchar(100);not null"` //标签名字
}

//实现gorm的接口，重命名表名
func (Tag) TableName() string {
	return "tag"
}
