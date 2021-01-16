# Go程序编译之后会得到一个可执行的二进制文件，其实在最终的镜像中是不需要go编译器的，也就是说我们只需要一个运行最终二进制文件的容器即可。
# 作为别名为"builder"的编译镜像，下面会用到
FROM golang AS builder

# 为镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 设置工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件app
RUN go build -o go-blog-app .

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 设置程序运行时必要的环境变量，包括监听端口、数据库配置等等
ENV SERVER_PORT=8085 \
    DATASOURCE_DRIVERNAME=mysql \
    DATASOURCE_HOST=127.0.0.1 \
    DATASOURCE_PORT=3306 \
    DATASOURCE_DATABASE=blog \
    DATASOURCE_USERNAME=root \
    DATASOURCE_PASSWORD=5KvA82*Ziq \
    DATASOURCE_CHARSET=utf8


# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/go-blog-app /

# 声明服务端口
EXPOSE 8085

# 启动容器时运行的命令
ENTRYPOINT ["/go-blog-app"]
