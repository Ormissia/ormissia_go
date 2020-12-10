// @File: blog
// @Date: 2020/12/6 20:54
// @Author: 安红豆
// @Description: 博客
package models

import "gorm.io/gorm"

//博客属性
type Blog struct {
	gorm.Model
	blogId      string //博客id
	User        User   //博客所属用户
	Type        Type   //类型
	Tags        []Tag  //标签
	TopImage    string //首图
	Title       string //标题
	Description string //描述
	Content     string //内容
	Visits      int    //访客数量
	IsDeleted   bool   //是否删除（正常、删除）
	IsRecommend bool   //是否推荐（推荐、普通）
	IsPublished bool   //是否发布（发布、草稿）
}
