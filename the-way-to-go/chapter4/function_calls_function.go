
// GOG

package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	// 注意, 这里 f2 定义的位置的作用域获得的 a 是全局的
	// 而不是 f1 中调用时使用 f1 的 a
	// 这个值在编译时已经决定好了
	print(a)
}
