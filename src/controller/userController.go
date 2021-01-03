// @File: userController
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 用户相关的控制器
package controller

import (
	"github.com/Ormissia/ormissia_go/src/common"
	"github.com/Ormissia/ormissia_go/src/dao"
	"github.com/Ormissia/ormissia_go/src/model"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (w *UserController) Register(ctx *gin.Context) {
	requestUser := model.User{}
	//从请求中获取用户参数
	err := ctx.Bind(&requestUser)
	if err != nil {
		util.ParameterError(ctx, err.Error(), nil)
		return
	}
	//获取参数
	username := requestUser.Username
	password := requestUser.Password
	email := requestUser.Email

	//用户名密码
	if len(username) == 0 || len(password) == 0 {
		util.ParameterError(ctx, util.GetCodeMsg(util.ErrorIllegalInput), nil)
		return
	}
	//邮箱格式验证(当邮箱长度不为0时进行格式验证)
	if len(email) != 0 && !util.StringRegexp(util.RegexpExpressionEmail, email) {
		util.ParameterError(ctx, util.GetCodeMsg(util.ErrorIllegalInput), nil)
		return
	}
	//判断用户是否存在
	existUser, _ := dao.SelectUserInfoByUsername(username)
	if existUser.ID != 0 {
		util.Error(ctx, util.GetCodeMsg(util.ErrorUsernameUsed), nil)
		return
	}

	//往数据库中插入该注册用户
	err = dao.InsertUser(requestUser)
	if err != nil {
		//插入出错
		util.Error(ctx, err.Error(), nil)
		return
	}
	//插入成功
	util.Success(ctx, util.GetCodeMsg(util.SuccessUserRegister), nil)
	return
}

//登录
func (w *UserController) Login(ctx *gin.Context) {
	requestUser := model.User{}
	//Bind在请求过程中，如果参数错误会直接抛异常返回400状态
	err := ctx.Bind(&requestUser)
	//ShouldBind在请求过程中，对参数检测不做处理
	//ctx.ShouldBind(&requestUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := dao.SelectUserInfoByUsername(requestUser.Username)

	//判断用户名密码是否正确
	if util.StringConvertMD5(requestUser.Password) != user.Password {
		ctx.JSON(http.StatusOK, gin.H{
			// 登录失败返回code 1001
			"code":    util.ErrorPasswordWrong,
			"message": util.GetCodeMsg(util.ErrorPasswordWrong),
		})
		return
	}

	//生成token
	token, _ := common.ReleaseToken(requestUser)
	//返回登录成功
	util.Response(ctx, http.StatusOK, util.HttpSuccess, "登录成功", nil, token)
	return
}
