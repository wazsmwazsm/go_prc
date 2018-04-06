package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Employee struct {
	Human // 嵌入字段 Human, Employee 会默认拥有 Human 的字段
	speciality string
	phone string // 和嵌入字段的属性重名的话, 以外部的为准
}

func main() {
	Bob := Employee{Human{"Bob", 34, "888-3333"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问 Human 的 phone 字段
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
