package main

import "fmt"

func main() {
	s1 := []byte{1, 2, 4, 5}
	fmt.Println(len(s1), cap(s1)) // 4 4

	s1 = Append(s1, []byte{5, 1}...)
	fmt.Println(s1, len(s1), cap(s1)) // [1 2 4 5 5 1] 6 12

	s1 = Append(s1, []byte{12, 44, 3, 2, 7}...)
	fmt.Println(s1, len(s1), cap(s1)) // [1 2 4 5 5 1 12 44 3 2 7] 11 12

	s1 = Append(s1, []byte{7, 8, 12}...)
	fmt.Println(s1, len(s1), cap(s1)) // [1 2 4 5 5 1 12 44 3 2 7 7 8 12] 14 28
}

func Append(slice []byte, data ...byte) []byte {
	length := len(slice)
	newLength := length + len(data)

	if newLength > cap(slice) { // 容量不够，扩容
		newSlice := make([]byte, newLength * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	// 容量够，直接切片
	slice = slice[0:newLength]
	copy(slice[length:newLength], data)

	return slice
}
