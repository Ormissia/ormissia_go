// @File: main
// @Date: 2020/12/6 20:47
// @Author: 安红豆
// @Description: 程序入口
package main

import (
	"github.com/ormissia/go-gin-blog/common"
	"github.com/ormissia/go-gin-blog/database"
	"github.com/ormissia/go-gin-blog/routers"
)

func main() {

	//初始化项目配置
	common.InitApplicationConfig()
	//初始化数据库
	database.InitMySql()
	//初始化路由并运行
	routers.InitRouter()
}
