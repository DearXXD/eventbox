# eventbox

事件中心, 负责事件收录


## 架构图


## 快速开发
make脚手架
```sh
➜  eventbox git:(master) ✗ make help
dep                            Get the dependencies
lint                           Lint Golang files
vet                            Run go vet
test                           Run unittests
test-coverage                  Run tests with coverage
build                          Local build
linux                          Linux build
run                            Run Server
clean                          Remove previous build
help                           Display this help screen
```

1. 使用go mod下载项目依赖
```sh
$ make dep
```

2. 添加配置文件(默认读取位置: etc/eventbox.toml)
```sh
$ 编辑样例配置文件 etc/eventbox.toml.example
$ mv etc/eventbox.toml.example etc/eventbox.toml
```

3. 启动服务
```sh
$ make run
```

## 相关文档