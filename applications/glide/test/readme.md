## glide 管理 go 依赖 （基于 vendor 文件夹）

## 安装 glide

go get github.com/Masterminds/glide

go install github.com/Masterminds/glide

## 安装依赖包

glide init 整理当前项目的依赖关系，生成 glide.yaml
	
glide install 安装依赖到 vendor 文件夹，生成 glide.lock
	
## 编译

到 main 包的目录下 go install

## 使用本地包时的问题

修改 glide.yaml, 忽略掉这些包
```yaml
ignore:
- youpackage1
- youpackage2

```
