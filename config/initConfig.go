// @File: blogController
// @Date: 2020/12/7 8:58
// @Author: 安红豆
// @Description: 初始化配置文件实例
package config

import (
	"github.com/spf13/viper"
	"ormissia_go/util"
	"os"
)

//初始化配置文件
func InitApplicationConfig(configFileName string) {

	workDir, _ := os.Getwd()
	viper.AddConfigPath(workDir)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(util.ConfigFileType)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
