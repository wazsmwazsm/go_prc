package main

import "fmt"

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	s2 := []byte{'b', 'o', 'b'}

	s3 := InsertStringSlice(s1, s2, 2)
	fmt.Printf("%c", s3)
}

func InsertStringSlice(slice, data []byte, pos int) []byte {
	
	slice1, slice2 := slice[:pos], slice[pos:]

	newSlice := make([]byte, len(slice1)) // 给足 newSlice 长度方便 copy
	copy(newSlice, slice1) // 这里要 copy，不要继续操作原 slice 的数组，防止数据无法被回收
	newSlice = append(newSlice, data...) // append 不像 copy，会自动扩容
	newSlice = append(newSlice, slice2...)
	return newSlice
}
