# 我的博客后端项目（Go语言版）
![Project Label](https://badgen.net/badge/github/Blog/blue?label=Golang)
[![Go Reference](https://pkg.go.dev/badge/Ormissia/ormissia_go.svg)](https://pkg.go.dev/github.com/ormissia/go-gin-blog)
[![Go Report Card](https://goreportcard.com/badge/github.com/ormissia/go-gin-blog)](https://goreportcard.com/report/github.com/ormissia/go-gin-blog)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/ormissia/go-gin-blog)](https://www.tickgit.com/browse?repo=github.com/ormissia/go-gin-blog)

> 为了功能上线方便，采用了Docker的方式部署程序

### [(提交代码——>功能自动发布)自动化部署流程(链接)](http://ormissia.com:13880/#/articleDetail/10)

#### 本地打包Docker镜像
进入工程目录后执行`docker build -t <image name> .`，当然要先安装Docker环境才行

Docker镜像打包完成后执行`docker run -d -p 8085:8085 <image name>`将程序跑起来

#### 本地开发配置
本地用IDE调试的时候需要在环境变量中添加相应的配置
```Environment
SERVER_PORT=8085 //访问端口
DATASOURCE_DRIVERNAME=mysql //数据库驱动类型
DATASOURCE_HOST=192.168.13.110 //数据库IP
DATASOURCE_PORT=3306 //数据库端口
DATASOURCE_DATABASE=blog //数据库名称
DATASOURCE_USERNAME=root //连接数据库用户名
DATASOURCE_PASSWORD=5KvA82*Ziq //连接数据库密码
DATASOURCE_CHARSET=utf8mb4 //连接数据库编码
```
GoLand/IDEA配置环境变量方式
![](https://ormissia-blog.oss-cn-qingdao.aliyuncs.com/image-hosting/GoLandEnvironment.jpg)
