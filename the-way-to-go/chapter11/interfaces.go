package main 

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1 // 实现了该接口，就可以赋值给该接口的接口变量

	fmt.Printf("The square has area: %f\n", areaIntf.Area())
} 
