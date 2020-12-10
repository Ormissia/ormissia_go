// @File: tag
// @Date: 2020/12/10 16:40
// @Author: 安红豆
// @Description: 标签
package models

import "gorm.io/gorm"

//标签属性
type Tag struct {
	gorm.Model
	TagId   string //标签id
	TagName string //标签名字
	Blogs   []Blog //标签所有博客
}
