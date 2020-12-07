// @File: blogController
// @Date: 2020/12/7 8:58
// @Author: 安红豆
// @Description: 初始化配置文件实例
package util

import (
	"github.com/spf13/viper"
)

//初始化配置文件
func InitApplicationConfig(configFilePath, configFileName string) {

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(ConfigFileType)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
