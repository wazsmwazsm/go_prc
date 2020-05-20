package main

// 指针逃逸
// go run -gcflags -m test1.go 查看逃逸分析
func main() {
	test()
}

func test() *int {
	a := 1

	return &a
}
