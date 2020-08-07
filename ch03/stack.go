package ch03

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Peek() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	top := len(*s) - 1

	return (*s)[top], true
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	top := len(*s) - 1
	result := (*s)[top]
	*s = (*s)[:top]

	return result, true
}

func (s *Stack) Len() int {
	return len(*s)
}
