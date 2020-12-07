// @File: initMySql_test
// @Date: 2020/12/6 22:04
// @Author: 安红豆
// @Description: 数据库连接的测试类
package database

import (
	"fmt"
	"testing"
)

func TestInitMySql(t *testing.T) {
	db := InitMySql()
	if db != nil {
		fmt.Println("Connect Successful!!!")
	} else {
		t.Errorf("Connect Filed!!!")
	}
}
