package ch03

// An unbound stack implemented using a slice.
// Values are stored in the slice in arrival order.
//
// The zero value and the nil value are safe to use:
//
//     // this is OK
//     var stack ch03.Stack = nil
//     stack.Push(42)
//
//     // this is also OK
//     stack := ch03.Stack{}
//     stack.Push(42)
type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Peek() (int, bool) {
	if s == nil || len(*s) == 0 {
		return 0, false
	}

	top := len(*s) - 1

	return (*s)[top], true
}

func (s *Stack) Pop() (int, bool) {
	if s == nil || len(*s) == 0 {
		return 0, false
	}

	top := len(*s) - 1
	result := (*s)[top]
	*s = (*s)[:top]

	return result, true
}

func (s *Stack) Len() int {
	if s == nil {
		return 0
	}

	return len(*s)
}

// Dump returns an slice of ints with the contents of the stack by order
// of arrival.
func (s *Stack) Dump() []int {
	result := make([]int, s.Len())
	copy(result, *s)

	return result
}
