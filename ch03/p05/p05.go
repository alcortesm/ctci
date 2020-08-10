package p05

import (
	"github.com/alcortesm/ctci/ch03"
)

type Stack struct {
	ch03.Stack
}

func (s *Stack) Sort() {
	sorted := ch03.Stack(make([]int, 0, len(s.Stack)))

	// move every element e in the stack to sorted
	for {
		e, ok := s.Stack.Pop()
		if !ok {
			break
		}

		// but before each move, move all elements in sorted that are
		// smaller than e to s.stack; this should result in all elements
		// in sorted being sorted with the smallest at the top.
		for {
			top, ok := sorted.Peek()
			if !ok || top >= e {
				break
			}

			_, _ = sorted.Pop()
			s.Stack.Push(top)
		}

		sorted.Push(e)
	}

	s.Stack = sorted
}
