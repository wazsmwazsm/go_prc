package main
import "fmt"
func main() {
	s := "good bye"
	// 和 C 不一样, go 的字符串是字面值, 不是指向第一个字符的指针
	// 所以这里用字符串指针可以直接修改字符串值
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
}
