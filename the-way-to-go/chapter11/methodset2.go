// Go 语言规范定义了接口方法集的调用规则：

// 类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
// 类型 T 的可调用方法集包含接受者为 T 的所有方法
// 类型 T 的可调用方法集不包含接受者为 *T 的方法

package main

import (
	"fmt"
)

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(v int) {
	*l = append(*l, v)
}

func CountInto(a Appender, start, end int) {
	for i := start; i < end; i++ {
		a.Append(i)
	}
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
} 

func main() {
	var lst List

	// CountInto(lst, 1, 10) // 报错, List 值没有实现 Appender
	if LongEnough(lst) { //  List 值实现了 Lener
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) // 指针变量实现了 Appender
	if LongEnough(plst) { // 指针变量实现了 Lener (其值实现了 Lener)
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
