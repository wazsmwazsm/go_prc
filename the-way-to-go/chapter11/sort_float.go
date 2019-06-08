package main

import (
	"fmt"
	"./packages/sort"
	"./packages/float"
)

func main() {
	float64()
}

func float64() {
	var f float.Float64Array
	f.Fill(10)
	sort.Sort(f)
	if ! sort.IsSorted(f) {
		panic("sort faild!")
	}

	fmt.Printf("%v", f)
}
