package main

// 闭包引用逃逸，static 变量实现的原理
// go run -gcflags -m test1.go 查看逃逸分析
func main() {
	f := test()

	for i := 0; i < 10; i++ {
		f()
	}
}

func test() func() int {
	i := 1

	return func() int {
		i++
		return i
	}
}
