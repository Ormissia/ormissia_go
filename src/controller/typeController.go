// @File: typeController
// @Date: 2020/12/23 10:00
// @Author: 安红豆
// @Description: 类型的控制层
package controller

import (
	"github.com/Ormissia/ormissia_go/src/dao"
	"github.com/Ormissia/ormissia_go/src/model"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
)

type TypeController struct {
}

//保存文章（新增和修改）
func (w *TypeController) SaveType(ctx *gin.Context) {
	articleType := model.Type{}
	//从请求中获取参数
	err := ctx.Bind(&articleType)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}

	//执行插入或修改操作
	if articleType.ID == 0 {
		//当id等于0时为新增操作
		err = dao.InsertType(articleType)
	} else {
		//当id不等0时为修改操作
		err = dao.UpdateType(articleType)
	}
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	util.Success(ctx, "保存成功", nil)
}

//删除文章
func (w *TypeController) DeleteType(ctx *gin.Context) {

}

//根据id查询文章
func (w *TypeController) SelectTypeById(ctx *gin.Context) {
	//从请求中获取要查询的文章id
	typeId := ctx.Query("id")
	articleType, err := dao.SelectTypeById(typeId)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	util.Success(ctx, util.GetCodeMsg(util.HttpSuccess), articleType)
	return
}

//根据分页参数查询类型
func (w *TypeController) SelectTypeByPage(ctx *gin.Context) {
	page := model.TypePage{}
	//从请求中获取参数
	err := ctx.Bind(&page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//根据分页参数查询类型列表
	articleTypes, err := dao.SelectTypeByPage(page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//根据分页参数查询博客总数
	count, err := dao.CountTypeByPage(page)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//返回成功
	util.Success(ctx, util.GetCodeMsg(util.HttpSuccess), map[string]interface{}{
		"count":    count,
		"dataList": articleTypes,
	})
	return
}
