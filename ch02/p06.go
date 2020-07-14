package ch02

import "fmt"

func IsPalindrome(l LinkedList) bool {
	hf, err := newHalfFinder(l)
	if err != nil {
		return true
	}

	// a simple stack using the same LinkedList type
	stack := LinkedList{}

	for n := l.First; n != nil; n = n.Next {
		switch {
		case hf.BeforeHalf():
			// push current element
			stack.First = &Node{
				Data: n.Data,
				Next: stack.First,
			}
			hf.Next()
		case hf.AtCenter():
			// don't store the center element in the stack, it is
			// irrelevant
			hf.Next()
		default: // we are now at the second half of the list
			// check each element agains the top of the stack
			if n.Data != stack.First.Data {
				return false
			}
			// pop stack
			stack.First = stack.First.Next
		}
	}

	return true
}

// a halfFinder knows whe you are at the first half, middle or at the
// second half of a linked list.
type halfFinder struct {
	slow     *Node
	fast     *Node
	atCenter bool
}

func newHalfFinder(l LinkedList) (*halfFinder, error) {
	if l.First == nil {
		return nil, fmt.Errorf("empty")
	}

	if l.First.Next == nil {
		return nil, fmt.Errorf("only 1 element")
	}

	return &halfFinder{
		slow: l.First,
		fast: l.First.Next,
	}, nil
}

func (h *halfFinder) Next() {
	if h.slow != nil {
		h.slow = h.slow.Next
	}

	h.atCenter = h.fast != nil &&
		h.fast.Next != nil &&
		h.fast.Next.Next == nil

	if h.fast != nil {
		if h.fast = h.fast.Next; h.fast != nil {
			h.fast = h.fast.Next
		}
	}
}

// first half doesn't include the center
func (h *halfFinder) BeforeHalf() bool {
	return h.fast != nil
}

func (h *halfFinder) AtCenter() bool {
	return h.atCenter
}
