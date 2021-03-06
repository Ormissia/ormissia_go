// @File: tagController
// @Date: 2020/12/23 10:00
// @Author: 安红豆
// @Description: 标签的控制层

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ormissia/go-gin-blog/dao"
	"github.com/ormissia/go-gin-blog/model"
	"github.com/ormissia/go-gin-blog/util"
)

type TagController struct {
}

//保存标签（新增和修改）
func (w *TagController) SaveTag(ctx *gin.Context) {
	articleTag := model.Tag{}
	//从请求中获取参数
	err := ctx.Bind(&articleTag)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}

	//执行插入或修改操作
	if articleTag.ID == 0 {
		//当id等于0时为新增操作
		//新增时将删除标记设为2，即为未删除
		articleTag.IsDeleted = 2
		err = dao.InsertTag(articleTag)
	} else {
		//当id不等0时为修改操作
		err = dao.UpdateTag(articleTag)
	}
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	util.Success(ctx, "保存成功", nil)
}

//删除标签
func (w *TagController) DeleteTag(ctx *gin.Context) {

}

//根据id查询文章
func (w *TagController) SelectTagById(ctx *gin.Context) {
	//从请求中获取要查询的文章id
	typeId := ctx.Query("id")
	articleTag, err := dao.SelectTagById(typeId)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	util.Success(ctx, util.GetCodeMsg(util.HttpSuccess), articleTag)
	return
}

//根据分页参数查询标签
func (w *TagController) SelectTagByPage(ctx *gin.Context) {
	page := model.TagPage{}
	//从请求中获取参数
	err := ctx.Bind(&page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//根据分页参数查询类型列表
	articleTags, err := dao.SelectTagByPage(page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//根据分页参数查询博客总数
	total, err := dao.CountTagByPage(page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//返回成功
	util.Success(ctx, util.GetCodeMsg(util.HttpSuccess), map[string]interface{}{
		"total":    total,
		"dataList": articleTags,
	})
	return
}
