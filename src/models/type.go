// @File: type
// @Date: 2020/12/10 16:40
// @Author: 安红豆
// @Description: 类型
package models

import "gorm.io/gorm"

//类型属性
type Type struct {
	gorm.Model
	TypeId   string //类型id
	TypeName string //类型名字
	Blogs    []Blog //类型下所有博客
}
