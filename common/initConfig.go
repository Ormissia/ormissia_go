// @File: initConfig
// @Date: 2020/12/7 8:58
// @Author: 安红豆
// @Description: 初始化配置文件实例
package common

import (
	"flag"
	"github.com/Ormissia/ormissia_go/util"
	"github.com/spf13/viper"
)

//初始化配置文件
func InitApplicationConfig(configFilePath, configFileName string) {
	//读取命令行配置文件的参数
	suffix := flag.String("config", "", "Config file suffix")
	flag.Parse()

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configFileName + *suffix)
	viper.SetConfigType(util.ConfigFileType)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
