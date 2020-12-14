// @File: initMySql_test
// @Date: 2020/12/6 22:04
// @Author: 安红豆
// @Description: 数据库连接的测试类
package database

import (
	"fmt"
	"github.com/Ormissia/ormissia_go/src/common"
	"github.com/Ormissia/ormissia_go/src/util"
	"testing"
)

func TestInitMySql(t *testing.T) {

	path := "../../"
	//初始化配置文件
	common.InitApplicationConfig(path, util.ConfigFileName)

	InitMySql()
	db := DB
	if db != nil {
		fmt.Println("Connect Successful!!!")
	} else {
		t.Errorf("Connect Filed!!!")
	}
}
