package main

import (
	"fmt"
	"strconv"
)
type T struct {
    a int
    b float32
    c string
}

func (t *T) String() string {
	return strconv.Itoa(t.a) + " / " + strconv.FormatFloat(float64(t.b), 'f', 6, 32) + " / " + t.c
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}

	fmt.Printf("%v\n", t)
}
