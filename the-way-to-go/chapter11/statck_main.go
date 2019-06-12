package main

import (
	"fmt"
	"./packages/stack"
)

func main() {
	// init
	st := stack.NewStack()
	fmt.Printf("stack is %v, empty?: %t, length: %v, top item: %v\n", st, st.IsEmpty(), st.Len(), st.Top())
	// push
	st.Push(12)
	st.Push("fuck")
	st.Push("5.67")
	fmt.Printf("stack is %v, empty?: %t, length: %v, top item: %v\n", st, st.IsEmpty(), st.Len(), (st.Top()))
	// pop
	if oldTop, err := st.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v poped!\n", oldTop)
		fmt.Printf("stack is %v, empty?: %t, length: %v, top item: %v\n", st, st.IsEmpty(), st.Len(), (st.Top()))
	}
	st.Pop()
	st.Pop()
	// empty statck
	if oldTop, err := st.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v poped!\n", oldTop)
		fmt.Printf("stack is %v, empty?: %t, length: %v, top item: %v\n", st, st.IsEmpty(), st.Len(), (st.Top()))
	}
	fmt.Printf("stack is %v, empty?: %t, length: %v, top item: %v\n", st, st.IsEmpty(), st.Len(), (st.Top()))
}
