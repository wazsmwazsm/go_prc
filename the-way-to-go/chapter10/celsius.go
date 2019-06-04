package main

import (
	"fmt"
	"strconv"
)

type Celsius float64

func (c *Celsius) String() string {
	return strconv.FormatFloat(float64(*c), 'f', 1, 64) + "°C"
}

func main() {
	var c Celsius = 23.5
	cp := &c // 注意，因为 String 方法是指针接收者, 只有指针实现了 Stringer 接口, 值没有

	fmt.Printf("%T %v\n", c)  // 值没有实现 Stringer 接口
	fmt.Printf("%T %v\n", cp) // 指针实现了
}
