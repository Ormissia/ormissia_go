// @File: main
// @Date: 2020/12/6 20:47
// @Author: 安红豆
// @Description: 程序入口
package main

import (
	"github.com/Ormissia/ormissia_go/src/database"
	"github.com/Ormissia/ormissia_go/src/routers"
	"github.com/Ormissia/ormissia_go/src/util"
)

func main() {

	//初始化配置文件
	util.InitApplicationConfig(util.ConfigFilePath, util.ConfigFileName)
	//初始化数据库
	database.InitMySql()
	//初始化路由并运行
	routers.InitRouter()
}
