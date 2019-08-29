## 使用 go mod 
1. 环境变量GOPATH不再用于解析imports包路径，即原有的GOPATH/src/下的包，通过import是找不到了。
2. Go Module功能开启后，下载的包将存放与$GOPATH/pkg/mod路径
3. $GOPATH/bin路径的功能依旧保持

## 环境变量设置（vendor 形式无需）
GOPATH : 下面的 src（go mod 不用这个）、bin（go install 生成的可执行文件）、pkg（pkg/mod 存放 go install 下载的依赖）
GOBIN : 设置为 $(GOPATH)/bin 就行

## 构建过程

1. 创建 main.go
2. 运行 go mod init test 生成 go.mod 文件, 指定 module 名称
3. 运行 go mod tidy 生成 require 依赖信息，生成 go.sum 文件
4. go install，运行 $(GOBIN)/test

## vendor 形式构建过程

1. 创建 main.go
2. 运行 go mod init test 生成 go.mod 文件, 指定 module 名称
3. 运行 go mod tidy 生成 require 依赖信息，生成 go.sum 文件
4. go mod vendor 生成 vendor 目录, 将项目的依赖下载到 $(GOPATH)/pkg/mod 中再移动到 vendor 中
5. go build -mod=vendor 生成可执行文件到当前目录


## 本地包使用
1. 创建本地宝，如本例的 greetings
2. 在包里 go mod init greetings 生成 go.mod 文件
3. 在包里 go mod tidy 梳理包的自身依赖, 生成 go.sum 文件
4. 修改项目 go.mod 添加 replace 或者使用 go mod edit -replace 命令来替换 github.com/wazsmwazsm/greetings（main 中的引用） 到本地路径 ./greetings
5. go install
