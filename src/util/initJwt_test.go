// @File: initJwt_test
// @Date: 2020/12/12 9:16
// @Author: 安红豆
// @Description: token生成的测试类
package util

import (
	"fmt"
	"github.com/Ormissia/ormissia_go/src/models"
	"testing"
)

func TestReleaseToken(t *testing.T) {
	user := models.User{
		UserId:   "123",
		Username: "ormissia",
	}

	tokenStr, err := ReleaseToken(user)
	if err == nil {
		fmt.Println("Successful ReleaseToken :" + tokenStr)
	} else {
		t.Errorf("ReleaseToken Failed :%s", err)
	}
}
