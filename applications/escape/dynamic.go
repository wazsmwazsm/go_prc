package main

// 动态类型逃逸（分析 interface{} 的内部结构，属于指针逃逸）
// go run -gcflags -m test1.go 查看逃逸分析
func main() {
	test()
}

func test() interface{} {
	i := 0
	return i
}
