package ch03

import "fmt"

// A limited capacity stack that efficiently allows to remove the bottom
// element.
type LinkedStack struct {
	top, bottom *node
	cap         int
	len         int
}

type node struct {
	value        int
	above, below *node
}

func NewLinkedStack(cap int) (*LinkedStack, error) {
	if cap < 1 {
		return nil, fmt.Errorf("invalid capacity")
	}

	return &LinkedStack{cap: cap}, nil
}

func (s *LinkedStack) Len() int {
	return s.len
}

func (s *LinkedStack) Dump() []int {
	result := make([]int, 0, s.len)

	for n := s.bottom; n != nil; n = n.above {
		result = append(result, n.value)
	}

	return result
}

func (s *LinkedStack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}

	return s.top.value, true
}

func (s *LinkedStack) Pop() (int, bool) {
	if s.len == 0 {
		return 0, false
	}

	toRemove := s.top

	if s.len == 1 {
		s.top = nil
		s.bottom = nil
	} else {
		s.top = toRemove.below
		s.top.above = nil
	}

	toRemove.below = nil // helps the Gargbage Collector
	s.len--

	return toRemove.value, true
}

func (s *LinkedStack) RemoveBottom() (int, bool) {
	if s.len == 0 {
		return 0, false
	}

	if s.len == 1 {
		return s.Pop()
	}

	toRemove := s.bottom
	s.bottom = toRemove.above
	s.bottom.below = nil
	toRemove.above = nil // helps the Garbage Collector
	s.len--

	return toRemove.value, true
}

func (s *LinkedStack) Push(v int) bool {
	if s.len == s.cap {
		return false
	}

	n := node{
		value: v,
	}

	if s.len == 0 {
		s.bottom = &n
		s.top = &n
	} else {
		n.below = s.top
		s.top.above = &n
		s.top = &n
	}

	s.len++

	return true
}
