// @File: userController
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 用户相关的控制器
package controller

import (
	"github.com/Ormissia/ormissia_go/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (w *UserController) Login(ctx *gin.Context) {
	requestUser := models.User{}
	//Bind在请求过程中，如果参数错误会直接抛异常返回400状态
	err := ctx.Bind(&requestUser)
	//ShouldBind在请求过程中，对参数检测不做处理
	//ctx.ShouldBind(&requestUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 没有数据库的用下面这个方法：这里先写死账号和密码  有数据库的要从数据库中获取
	if requestUser.Username != "123" || requestUser.Password != "123" {
		ctx.JSON(http.StatusOK, gin.H{
			// 登录失败返回code 1001
			"code":    1001,
			"message": "failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		// 登录失败返回code 1000
		"code":    1000,
		"message": "success",
	})
	return
}
