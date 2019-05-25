package main

import (
	"fmt"
)

func main() {
	k := 6
	switch k {
	case 4: fmt.Println("was <= 4"); fallthrough;
	case 5: fmt.Println("was <= 5"); fallthrough;
	// fallthrough 会不管后面的条件是否成立, 会强行向下执行
	// 下面的条件都会被执行, 这点和 C 是不同的
	case 6: fmt.Println("was <= 6"); fallthrough;
	case 7: fmt.Println("was <= 7"); fallthrough;
	case 8: fmt.Println("was <= 8"); fallthrough;
	default: fmt.Println("default case")
	}
}
