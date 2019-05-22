package main

import (
	"fmt"
	"container/list"
)

func main() {
	l := list.New()

	l.PushBack(102)
	l.PushBack(103)
	l.PushFront(101)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
