package p02

import (
	"fmt"

	"github.com/alcortesm/ctci/ch03"
)

type StackWithMin struct {
	// the stack keeping the values
	values *ch03.Stack
	// the stack keeping the indexes in "values" of the minimum values
	indexOfMin *ch03.Stack
}

func NewStackWithMin() *StackWithMin {
	return &StackWithMin{
		values:     &ch03.Stack{},
		indexOfMin: &ch03.Stack{},
	}
}

func (s StackWithMin) Push(v int) {
	s.values.Push(v)

	// if there was no min, just add the first element as the current
	// min
	min, ok := s.Min()
	if !ok {
		s.indexOfMin.Push(0)
		return
	}

	// if there was a min, compare with this new element and update if
	// necessary.
	if v < min {
		s.indexOfMin.Push(s.values.Len() - 1)
	}
}

func (s StackWithMin) Pop() (int, bool) {
	result, ok := s.values.Pop()
	if !ok {
		return 0, false
	}

	// if the current min was at the index of the element we have just
	// removed then we must remove that index.
	if i, _ := s.indexOfMin.Peek(); i == s.values.Len() {
		_, _ = s.indexOfMin.Pop()
	}

	return result, true
}

func (s StackWithMin) Peek() (int, bool) {
	return s.values.Peek()
}

func (s StackWithMin) Min() (int, bool) {
	i, ok := s.indexOfMin.Peek()
	if !ok {
		return 0, false
	}

	return (*s.values)[i], true
}

func (s StackWithMin) String() string {
	return fmt.Sprintf("%d %d", s.values, s.indexOfMin)
}
