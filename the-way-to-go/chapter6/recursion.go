package main

import "fmt"

func main() {
	pp(10)
}

// 递归打印 n ~ 1
func pp(n int ) {
	if n < 1 {
		return
	}
	
	fmt.Println(n)
	
	pp(n - 1)
}
