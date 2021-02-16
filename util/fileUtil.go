// @File: fileUtil
// @Date: 2020/12/17 14:03
// @Author: 安红豆
// @Description: 文件相关的工具方法

package util

import "os"

func PathExists(path string) (err error) {
	_, err = os.Stat(path)
	//如果返回值错误类型用os.IsNotExist判断为true说明文件或文件夹不存在
	if os.IsNotExist(err) {
		//创建路径
		err = os.MkdirAll(path, 664)
	}
	//创建后重新判断路径是否存在
	_, err = os.Stat(path)
	if err != nil {
		return err
	}
	//如果err为nil说明文件或文件夹存在
	return nil
}
