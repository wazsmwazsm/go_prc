package main

import (
	"fmt"
	"strings"
)

func main() {

	f := func(r rune) rune {
		// 7 位的 ASCII 码，不属于的则返回 ? 
		if r > 127 || r < 0 {
			return '?'
		}

		return r
	}

	fmt.Println(strings.Map(f, "hello 我亲爱的 friend"))
}

