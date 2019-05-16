package main

import (
	"fmt"
	"errors"
)

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}

	s2, err := RemoveStringSlice(s1, 1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%c", s2)
	}
	
}

func RemoveStringSlice(slice []byte, start, end int) ([]byte, error) {
	
	if start < 0 || start > end || end > len(slice) {
		return nil, errors.New("Out of range!")
	}

	slice1, slice2 := slice[:start], slice[end + 1:]

	newSlice := make([]byte, len(slice1)) // 给足 newSlice 长度方便 copy
	copy(newSlice, slice1) // 这里要 copy，不要继续操作原 slice 的数组，防止数据无法被回收
	newSlice = append(newSlice, slice2...)
	return newSlice, nil
}
