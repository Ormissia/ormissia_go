// @File: routers
// @Date: 2020/12/7 22:08
// @Author: 安红豆
// @Description: 统一注册路由
package routers

import (
	"github.com/gin-gonic/gin"
	"ormissia_go/src/middleware"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	//将中间件注册到全局，对所有路由生效
	r.Use(middleware.Cors())
	//注册路由
	LoginRouter(r)

	return r
}
