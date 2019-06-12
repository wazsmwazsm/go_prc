package stack

import (
	"errors"
	"strconv"
)
// 顺序栈
type stack struct {
	index int
	arr stackArr
}

func NewStack() *stack {
	s := new(stack)
	s.index = -1
	return s
}

func (s *stack) String() string {
	var str string
	str = "Stack head: " + strconv.Itoa(s.index) + " Stack Field: "
	for i := 0; i < LIMIT; i++ {
		str += "[" + strconv.Itoa(s.arr[i]) + "]"
	} 

	return str
}

func (s *stack) IsEmpty() bool {
	if s.index == -1 {
		return true
	}

	return false
}

func (s *stack) IsFull() bool {
	if s.index == LIMIT - 1 {
		return true
	}

	return false
}

func (s *stack) Pop() (int, error) {
	if ! s.IsEmpty() {
		popItem := s.arr[s.index]
		s.arr[s.index] = 0
		s.index--

		return popItem, nil
	}

	return 0, errors.New("Statck is empty!")
}

func (s *stack) Push(item int) error {
	if ! s.IsFull() {
		s.index++
		s.arr[s.index] = item

		return nil
	} 

	return errors.New("Statck is full!")
}
