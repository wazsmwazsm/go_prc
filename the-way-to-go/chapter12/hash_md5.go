package main

import (
	"fmt"
	"crypto/md5"
	"io"
)

func main() {
	hasher := md5.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
}