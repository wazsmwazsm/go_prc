package main

// 栈空间不足逃逸
// go run -gcflags -m test1.go 查看逃逸分析
func main() {
	test()
}

func test() {
	// s := make([]int, 10, 10) // 小空间占用，不逃逸
	s := make([]int, 1000000, 1000000) // 逃逸
	s[0] = 1
}
