// @File: initConfig
// @Date: 2020/12/7 8:58
// @Author: 安红豆
// @Description: 初始化配置文件实例

package common

import (
	"github.com/ormissia/go-gin-blog/util"
	"github.com/spf13/viper"
)

//初始化配置文件
func InitApplicationConfig() {

	if err := viper.BindEnv(util.ConfigServerPort); err != nil {
		panic(err)
	}
	if err := viper.BindEnv(util.ConfigDataSourceDriverName); err != nil {
		panic(err)
	}
	if err := viper.BindEnv(util.ConfigDataSourceHost); err != nil {
		panic(err)
	}
	if err := viper.BindEnv(util.ConfigDataSourcePort); err != nil {
		panic(err)
	}
	if err := viper.BindEnv(util.ConfigDataSourceDatabase); err != nil {
		panic(err)
	}
	if err := viper.BindEnv(util.ConfigDataSourceUsername); err != nil {
		panic(err)
	}
	if err := viper.BindEnv(util.ConfigDataSourcePassword); err != nil {
		panic(err)
	}
	if err := viper.BindEnv(util.ConfigDataSourceCharset); err != nil {
		panic(err)
	}
	if err := viper.BindEnv(util.ConfigUploadFilePathImage); err != nil {
		panic(err)
	}
}
