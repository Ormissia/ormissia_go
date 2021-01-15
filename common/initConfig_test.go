// @File: initConfig_test
// @Date: 2020/12/7 9:16
// @Author: 安红豆
// @Description: 初始化配置文件实例的测试类
package common

import (
	"github.com/Ormissia/ormissia_go/util"
	"github.com/spf13/viper"
	"testing"
)

func TestInitApplicationConfig(t *testing.T) {

	path := "../../"

	InitApplicationConfig(path, util.ConfigFileName)

	tests := []struct{ a, b string }{
		{util.ConfigServerPort, "8085"},
		{util.ConfigDataSourceDriverName, "mysql"},
		{util.ConfigDataSourceHost, "127.0.0.1"},
		{util.ConfigDataSourcePort, "3306"},
		{util.ConfigDataSourceDatabase, "blog"},
		{util.ConfigDataSourceUsername, "root"},
		{util.ConfigDataSourcePassword, "123"},
		{util.ConfigDataSourceCharset, "utf8"},
	}

	for _, tt := range tests {
		if actual := viper.GetString(tt.a); actual != tt.b {
			t.Errorf("viper.GetString(%s); got %s; expected %s;", tt.a, actual, tt.b)
		}
	}
}
