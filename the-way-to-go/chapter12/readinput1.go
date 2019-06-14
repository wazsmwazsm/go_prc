package main

import "fmt"

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter you full name: ")
	// fmt.Scanln(&firstName, &lastName) // 将变量绑定 stdin
	fmt.Scanf("%s %s", &firstName, &lastName) // 按照格式输入
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s) // 字符串做输入
	fmt.Println("From the string we read: ", f, i, s)
}
