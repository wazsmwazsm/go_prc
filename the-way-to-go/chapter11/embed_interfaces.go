package main

import "fmt"

type Shaper interface {
	Area() float32
	Perimeter() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return sq.side * 4
}

type MySquare struct {
	Shaper // 内嵌接口，MySquare 就实现了接口
	side   float32
	name   string
}

// 方法重载
func (sq *MySquare) Perimeter() float32 {
	return sq.side * 5
}

func main() {
	// 传入满足 Shaper 接口的类型
	sq := &MySquare{&Square{4}, 3, "MySquare is 5 side!"}
	fmt.Printf("mysquare: %v\nperimeter: %.2f, area: %.2f\n", sq, sq.Perimeter(), sq.Area())
}
