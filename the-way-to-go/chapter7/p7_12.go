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

	return append(slice[:start], slice[end + 1:]...), nil
}
