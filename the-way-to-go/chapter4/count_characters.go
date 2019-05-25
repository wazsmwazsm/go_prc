package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 统计字符串的字符数
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))

	str2 := "asSASA ddd dsjkdsjsこん dk"
	// len 数的是所占字符的数量
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	// utf8.RuneCountInString 数的是所有编码真实字符的数量
	fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))
}

/* Output:
The number of bytes in string str1 is 22
The number of characters in string str1 is 22
The number of bytes in string str2 is 28
The number of characters in string str2 is 24
*/
