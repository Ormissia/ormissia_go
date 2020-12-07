// @File: blogController
// @Date: 2020/12/7 9:16
// @Author: 安红豆
// @Description: 初始化配置文件实例的测试类
package util

import (
	"github.com/spf13/viper"
	"testing"
)

func TestInitApplicationConfig(t *testing.T) {

	path := "../config"

	InitApplicationConfig(path, ConfigFileName)

	tests := []struct{ a, b string }{
		{ConfigServerPort, "8085"},
		{ConfigDataSourceDriverName, "mysql"},
		{ConfigDataSourceHost, "127.0.0.1"},
		{ConfigDataSourcePort, "3306"},
		{ConfigDataSourceDatabase, "blog"},
		{ConfigDataSourceUsername, "root"},
		{ConfigDataSourcePassword, "123"},
		{ConfigDataSourceCharset, "utf8"},
	}

	for _, tt := range tests {
		if actual := viper.GetString(tt.a); actual != tt.b {
			t.Errorf("viper.GetString(%s); got %s; expected %s;", tt.a, actual, tt.b)
		}
	}
}
