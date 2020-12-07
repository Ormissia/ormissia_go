// @File: main
// @Date: 2020/12/6 20:47
// @Author: 安红豆
// @Description: 程序入口
package main

import (
	"ormissia_go/config"
	"ormissia_go/util"
)

func main() {

	//初始化配置文件
	config.InitApplicationConfig(util.ConfigFileName)
}
