package main

import (
	"fmt"
	"container/list"
)

func main() {
	l := list.New() // 创建双向链表
	e4 := l.PushBack(4) // 列表尾添加元素，值为 4
	e1 := l.PushFront(1) // 列表头添加元素，值为 1
	e3 := l.InsertBefore(3, e4) // 元素 e4 前面插入新元素，值为 3 
	e2 := l.InsertAfter(2, e1) // 元素 e1 后面插入新元素，值为 2
	fmt.Println(e3, e2, e3.Value, e2.Value)
	// 遍历链表
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
