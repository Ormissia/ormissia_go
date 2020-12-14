// @File: type
// @Date: 2020/12/10 16:40
// @Author: 安红豆
// @Description: 类型
package model

import "gorm.io/gorm"

//类型属性
type Type struct {
	Article []Article `json:"article" gorm:"-"` //类型下的文章
	gorm.Model
	TypeName string `json:"type_name" gorm:"type:varchar(20);not null"` //类型名字
}

//实现gorm的接口，重命名表名
func (Type) TableName() string {
	return "type"
}