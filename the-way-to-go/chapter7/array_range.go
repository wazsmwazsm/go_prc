package main

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}

	// item 只是 for range 的一个拷贝，修改不会影响原数组
	for _, item := range items {
		item *= 2
	}

	fmt.Println(items)

	for k := range items {
		items[k] *= 2
	}
	
	fmt.Println(items)
}
