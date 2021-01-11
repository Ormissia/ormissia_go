# Go程序编译之后会得到一个可执行的二进制文件，其实在最终的镜像中是不需要go编译器的，也就是说我们只需要一个运行最终二进制文件的容器即可。
# 作为别名为"builder"de编译镜像，下面会用到
FROM golang AS builder

MAINTAINER ormissia "ormissia@outlook.com"

# 为镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件app
RUN go build -o app .

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 将配置文件拷贝到根目录

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/app /

# 声明服务端口
EXPOSE 8085

# 启动容器时运行的命令
CMD ["/app"]
