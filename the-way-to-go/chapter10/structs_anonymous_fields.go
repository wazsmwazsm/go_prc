package main

import "fmt"

type innerS struct {
	in1, in2 int
}

type outerS struct {
	b int
	c float32
	int // 匿名字段
	innerS // 匿名字段，结构体嵌套
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 6.2
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	inner := innerS{99, 12}
	outer2 := outerS{6, 7.5, 60, inner}
	fmt.Println("outer2 is:", outer2)
	fmt.Println(outer2.in1, outer2.in2)
}
