// @File: routers
// @Date: 2020/12/7 22:08
// @Author: 安红豆
// @Description: 统一注册路由
package routers

import (
	"github.com/Ormissia/ormissia_go/src/middleware"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//初始化路由
func InitRouter() {
	//注册路由
	r := gin.Default()
	r = collectRoutes(r)

	//运行
	port := viper.GetString(util.ConfigServerPort)
	panic(r.Run(":" + port))
}

//注册所有路由
func collectRoutes(r *gin.Engine) *gin.Engine {

	//将中间件注册到全局，对所有路由生效
	r.Use(middleware.Cors(), middleware.ServerError())

	//注册路由
	UserRouter(r)
	ArticleRouter(r)
	FileRouter(r)
	TypeRouter(r)
	TagRouter(r)

	return r
}
