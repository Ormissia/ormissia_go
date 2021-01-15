// @File: regexpUtil_test
// @Date: 2020/12/7 9:16
// @Author: 安红豆
// @Description: 正则验证的测试类
package util

import "testing"

func TestEmailRegexp(t *testing.T) {
	tests := []struct {
		email   string
		matched bool
	}{
		{"ormissia@outlook.com", true},
		{"1432@qq.com", true},
		{"   @qq.com", false},
		{"", false},
		{"@outlook.com", false},
	}

	for _, tt := range tests {
		if matched := StringRegexp(RegexpExpressionEmail, tt.email); matched != tt.matched {
			t.Errorf("util.EmailRegexp(%s); got %t; expected %t;", tt.email, matched, tt.matched)
		}
	}
}
