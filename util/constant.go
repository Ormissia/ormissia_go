// @File: blogController
// @Date: 2020/12/7 9:01
// @Author: 安红豆
// @Description: 常量类
package util

const (
	//配置文件相关的常量
	ConfigFilePath    = "./config"        //文件路径
	ConfigFileName    = "application"     //文件名
	ConfigFileNameDev = "application-dev" //文件名
	ConfigFileType    = "yml"             //文件类型
)

const (
	//程序相关的配置
	ConfigServer     = "server."
	ConfigServerPort = ConfigServer + "port" //端口
)

const (
	//数据库配置相关的常量
	ConfigDataSource           = "datasource."
	ConfigDataSourceDriverName = ConfigDataSource + "driverName" //驱动名称
	ConfigDataSourceHost       = ConfigDataSource + "host"       //数据库IP
	ConfigDataSourcePort       = ConfigDataSource + "port"       //数据库端口
	ConfigDataSourceDatabase   = ConfigDataSource + "database"   //数据库名称
	ConfigDataSourceUsername   = ConfigDataSource + "username"   //数据库用户名
	ConfigDataSourcePassword   = ConfigDataSource + "password"   //数据库密码
	ConfigDataSourceCharset    = ConfigDataSource + "charset"    //数据库编码格式
)
