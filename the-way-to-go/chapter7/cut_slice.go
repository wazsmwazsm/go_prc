package main

import (
	"fmt"
	"errors"
)

func main() {
	s1 := []byte{1, 3, 4, 99, 8, 3, 6} 

	s2, s3, err := cut(s1, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s2, s3)
	}
}

func cut(s []byte, n int) ([]byte, []byte, error) {
	if n > cap(s) {
		return nil, nil, errors.New("Out of range!")
	}
	return s[:n], s[n:], nil
}
