package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	b.WriteString("aaa")
	b.WriteByte('b')
	b.Write([]byte{'c', 'd'})

	fmt.Println(b.String())
}
