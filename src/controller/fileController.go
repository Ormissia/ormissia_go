// @File: fileController
// @Date: 2020/12/17 9:24
// @Author: 安红豆
// @Description: 文件的路由配置
package controller

import (
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

type FileController struct {
}

//接收上传的图片
func (w *FileController) UploadImage(ctx *gin.Context) {
	//从请求中获取图片文件
	file, _ := ctx.FormFile("image")
	//获取文件名
	fileName := file.Filename
	//拼接年月路径
	yearMonth := time.Now().Format(util.DataFormatPath)
	//从配置文件中获取图片存放的路径并拼接年月路径
	path := viper.GetString(util.ConfigUploadFilePathImage) + yearMonth + "/"
	//判断路径是否存在，不存在则创建
	err := util.PathExists(path)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	//保存图片
	err = ctx.SaveUploadedFile(file, path+fileName)
	if err != nil {
		util.Error(ctx, err.Error(), nil)
		return
	}
	util.Success(ctx, "保存成功", nil)
}
