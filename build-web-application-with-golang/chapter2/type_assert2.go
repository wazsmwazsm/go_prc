package main

import (
	"fmt"
	"strconv"
)

type Element interface{} // 空接口
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func main() {
	list := make(List, 3)
	// 空接口默认被所有类型实现, 接口值可以保存任何类型的值
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for index, el := range list {
		// 使用 switch 选择类型
		// 注意，el 必须是空接口值才行
		switch value := el.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}
