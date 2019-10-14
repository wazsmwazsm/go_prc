package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(111)

	if s.items[0].(int) != 111 || s.GetLen() != 1 || s.GetCap() != 1 {
		t.Error("push err")
	}

	s.Push(222)
	if s.items[1].(int) != 222 || s.GetLen() != 2 || s.GetCap() != 2 {
		t.Error("push err")
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	s.Push(33)
	s.Push(44)

	if s.Pop().(int) != 44 || s.GetLen() != 1 || s.GetCap() != 2 {
		t.Error("push err")
	}

	if s.Pop().(int) != 33 || s.GetLen() != 0 || s.GetCap() != 2 {
		t.Error("push err")
	}
}

func BenchmarkPushPop(b *testing.B) {
	s := NewStack()
	n := b.N
	for i := 0; i < n; i++ {
		s.Push(i)
	}
	for i := 0; i < n; i++ {
		s.Pop()
	}
}
