package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a, b int
}

func (tn *TwoInts) String() string {
	// 注意，String 方法中不要调用调了 String 方法的方法。如 fmt.Println 之类
	// 会导致一直嵌套
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}
