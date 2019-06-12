// 实现一个栈结构
package stack

import "errors"

type Stacker interface {
	Len() int
	IsEmpty() bool
	Push(v interface{})
	Pop() (interface{}, error)
	Top() interface{}
}

type Item struct {
	data interface{}
	next *Item
}

func (this *Item) GetData() interface{} {
	return this.data
}

func (this *Item) GetNext() *Item {
	return this.next
}

func NewItem(data interface{}, next *Item) *Item {
	return &Item{data, next}
}

// 链式栈
type Stack struct {
	top *Item
	length int
}

func (this *Stack) Len() int {
	return this.length
}

func (this *Stack) IsEmpty() bool {
	if this.length == 0 {
		return true
	}
	return false
}

func (this *Stack) Push(v interface{}) {
	newTop := &Item{v, this.top}
	this.top = newTop
	this.length++
}

func (this *Stack) Pop() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("Stack is Empty!")
	}
	oldTop := this.top
	this.top = this.top.next
	this.length--
	return oldTop.data, nil
}

func (this *Stack) Top() interface{} {
	if this.top == nil {
		return nil
	}
	return this.top.data
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}
