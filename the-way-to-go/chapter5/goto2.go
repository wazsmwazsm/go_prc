// compile error goto2.go:8: goto TARGET jumps over declaration of b at goto2.go:8
package main

import "fmt"

func main() {
		a := 1
		goto TARGET // 编译失败, 标签和 goto 之间不能有定义新变量的语句
		b := 9
	TARGET:
		b += a
		fmt.Printf("a is %v *** b is %v", a, b)
}
