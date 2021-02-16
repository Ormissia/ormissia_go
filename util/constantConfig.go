// @File: constantConfig
// @Date: 2020/12/7 9:01
// @Author: 安红豆
// @Description: 配置文件相关的常量

package util

const (
	//程序相关的配置
	ConfigServer     = "SERVER_"
	ConfigServerPort = ConfigServer + "PORT" //端口
)

const (
	//数据库配置相关的常量
	ConfigDataSource           = "DATASOURCE_"
	ConfigDataSourceDriverName = ConfigDataSource + "DRIVER_NAME" //驱动名称
	ConfigDataSourceHost       = ConfigDataSource + "HOST"        //数据库IP
	ConfigDataSourcePort       = ConfigDataSource + "PORT"        //数据库端口
	ConfigDataSourceDatabase   = ConfigDataSource + "DATABASE"    //数据库名称
	ConfigDataSourceUsername   = ConfigDataSource + "USERNAME"    //数据库用户名
	ConfigDataSourcePassword   = ConfigDataSource + "PASSWORD"    //数据库密码
	ConfigDataSourceCharset    = ConfigDataSource + "CHARSET"     //数据库编码格式
)

const (
	//文件相关的配置
	ConfigFileUploadFile      = "filePath."
	ConfigUploadFilePathImage = ConfigFileUploadFile + "image" //图片路径
)
