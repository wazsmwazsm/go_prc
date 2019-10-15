package stack

import (
	"sync"
)

// Stack a slice stack
type Stack struct {
	items []interface{}
	sync.Mutex
}

// NewStack init a stack with a init cap
func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

// GetCap get stack capacity
func (s *Stack) GetCap() int {
	return cap(s.items)
}

// GetLen get stack length
func (s *Stack) GetLen() int {
	return len(s.items)
}

// Push push a item to stack
func (s *Stack) Push(item interface{}) {
	s.Lock()
	defer s.Unlock()

	s.items = append(s.items, item)
}

// Pop pop a item from stack
func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()
	length := s.GetLen()
	if length == 0 {
		return nil
	}

	item := s.items[length-1]
	s.items = s.items[:length-1]

	return item
}
