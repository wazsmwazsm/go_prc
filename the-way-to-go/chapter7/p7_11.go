package main

import (
	"fmt"
	"errors"
)

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	s2 := []byte{'b', 'o', 'b'}

	s3, err := InsertStringSlice(s1, s2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%c", s3)
	}
	
}

func InsertStringSlice(slice, data []byte, pos int) ([]byte, error){
	
	if pos < 0 || pos > len(slice) {
		return nil, errors.New("Out of range!")
	}

	return append(slice[:pos], append(data, slice[pos:]...)...), nil
	
}
