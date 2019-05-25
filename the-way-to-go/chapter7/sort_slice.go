package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{22, 4, 11, 2, 1, 7, 5}
	sort.Ints(s1)
	fmt.Println(s1)

	s2 := []string{"Jack", "Bob", "Mariy", "Alice"}
	sort.Strings(s2)
	fmt.Println(s2)

	if sort.IntsAreSorted(s1) {
		result := sort.SearchInts(s1, 1) // 返回值的位置
		fmt.Println(result)
		result := sort.SearchInts(s1, 88) // 搜不到的，返回 length
		fmt.Println(result)
	}
}