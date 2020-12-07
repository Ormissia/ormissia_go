// @File: userController
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 用户相关的控制器
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

type Login struct {
	Username string
	Password string
}

func (w *UserController) Login(c *gin.Context) {
	var login Login
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login = Login{
		Username: "123",
		Password: "123",
	}
	// 没有数据库的用下面这个方法：这里先写死账号和密码  有数据库的要从数据库中获取
	if login.Username != "123" || login.Password != "123" {
		c.JSON(http.StatusOK, gin.H{
			// 登录失败返回code 1001
			"code":    1001,
			"message": "failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		// 登录失败返回code 1000
		"code":    1000,
		"message": "success",
	})
	return
}
