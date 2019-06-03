package main

import "fmt"

type B struct {
	thing int
}

// 指针接收者
func (b *B) change() {
	b.thing = 1
}
// 值接受者, b 是一个拷贝
func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {
	var b1 B
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())
}
