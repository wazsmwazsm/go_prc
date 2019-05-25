package main

import "fmt"

func main() {
	a := []byte{'s', 'i', 'l', 'l', 'y'}
	b := []byte{'s', 'i', 'l', 'l', 'y'}
	c := []byte{'s', 'c', 'o', 'r', 'e'}
	d := []byte{'s', 'i', 'l', 'l', 'y', '!'}

	fmt.Println(compare(a, b))
	fmt.Println(compare(a, c))
	fmt.Println(compare(a, d))
}

func compare(a, b []byte) int {
	// 字符相等判断
	for i := 0; i < len(a) && i < len(b); i++ {
		switch  {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1				
		}
	}
	// 长度判断
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1	
	}

	return 0 // 相等
}
