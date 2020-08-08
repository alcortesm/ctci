package p03

import (
	"fmt"
	"log"

	"github.com/alcortesm/ctci/ch03"
)

type SetOfStacks struct {
	stacks []*ch03.LinkedStack
	cap    int // capacity of each stack
}

func NewSetOfStacks(cap int) (*SetOfStacks, error) {
	if cap < 1 {
		return nil, fmt.Errorf("invalid maximum size (%d)", cap)
	}

	return &SetOfStacks{
		cap: cap,
	}, nil
}

func (s *SetOfStacks) Dump() [][]int {
	result := [][]int{}

	for _, s := range s.stacks {
		result = append(result, s.Dump())
	}

	return result
}

func (s *SetOfStacks) Push(v int) {
	if len(s.stacks) == 0 {
		newStack, err := ch03.NewLinkedStack(s.cap)
		if err != nil {
			log.Fatal(err)
		}

		s.stacks = []*ch03.LinkedStack{newStack}
	}

	topStack := s.stacks[len(s.stacks)-1]

	if topStack.Len() == s.cap {
		var err error
		topStack, err = ch03.NewLinkedStack(s.cap)
		if err != nil {
			log.Fatal(err)
		}
		s.stacks = append(s.stacks, topStack)
	}

	topStack.Push(v)
}

func (s *SetOfStacks) Pop() (int, bool) {
	if len(s.stacks) == 0 {
		return 0, false
	}

	topStack := s.stacks[len(s.stacks)-1]

	result, ok := topStack.Pop()
	if !ok {
		return 0, false
	}

	if topStack.Len() == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	return result, true
}

func (s *SetOfStacks) Peek() (int, bool) {
	if len(s.stacks) == 0 {
		return 0, false
	}

	topStack := s.stacks[len(s.stacks)-1]
	return topStack.Peek()
}

func (s *SetOfStacks) Len() int {
	result := 0

	for _, s := range s.stacks {
		result += s.Len()
	}

	return result
}

func (s *SetOfStacks) PopAt(stackNum int) (int, bool) {
	if stackNum < 0 || stackNum >= len(s.stacks) {
		return 0, false
	}

	current := s.stacks[stackNum]
	result, ok := current.Pop()
	if !ok {
		return 0, false
	}

	// if there are more stacks, shift their elements down
	if stackNum < len(s.stacks)-1 {
		v, ok := s.shiftDown(stackNum + 1)
		if ok {
			if ok = current.Push(v); !ok {
				log.Fatal("unreachable: cannot push shifted value, it should have fitted though")
			}
		}
	}

	if current.Len() == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	return result, true
}

// shift the values in a stack and all their following stacks one
// position down, returning the shifted value if any. It deletes stacks
// if they become empty as a consequence of a shift down operation.
func (s *SetOfStacks) shiftDown(stackNum int) (int, bool) {
	if stackNum < 0 || stackNum >= len(s.stacks) {
		return 0, false
	}

	current := s.stacks[stackNum]

	result, ok := current.RemoveBottom()
	if !ok {
		return 0, false
	}

	// if there are more stacks after this one,
	// shift them down too
	if stackNum < len(s.stacks)-1 {
		v, ok := s.shiftDown(stackNum + 1)
		if ok {
			if ok = current.Push(v); !ok {
				log.Fatal("unreachable: cannot push shifted value, it should have fitted though")
			}
		}
	}

	// remove the current stack if it becomes empty
	// it only becomes empty if it is the last one
	// so we can safetly simplify the operation here
	// to remove the last stack.
	if current.Len() == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	return result, true
}
