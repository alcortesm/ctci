package ch03

import (
	"errors"
	"fmt"
)

type MultiStack struct {
	info   []info
	values []int
}

func NewMultiStack(nStacks, defaultSize int) (*MultiStack, error) {
	if nStacks < 3 {
		return nil, fmt.Errorf("invalid number of stacks: (%d)", nStacks)
	}

	if defaultSize < 1 {
		return nil, fmt.Errorf("invalid default size (%d)", defaultSize)
	}

	result := &MultiStack{
		values: make([]int, nStacks*defaultSize),
		info:   make([]info, nStacks),
	}

	for i := range result.info {
		result.info[i].start = i * defaultSize
		result.info[i].cap = defaultSize
	}

	return result, nil
}

func (ms *MultiStack) String() string {
	return fmt.Sprintf("%d", ms.values)
}

// PeekAll returns a copy of the contents of all the stacks.
func (ms *MultiStack) PeekAll() [][]int {
	result := make([][]int, len(ms.info))

	for i, info := range ms.info {
		result[i] = make([]int, info.len)
		for j := 0; j < info.len; j++ {
			index := ms.index(info.start + j)
			result[i][j] = ms.values[index]
		}
	}

	return result
}

var ErrFull = errors.New("full stack")

func (ms *MultiStack) Push(stackNum, value int) error {
	if stackNum >= len(ms.info) {
		return fmt.Errorf("invalid stack num (%d)", stackNum)
	}

	if ms.isFull() {
		return ErrFull
	}

	info := &ms.info[stackNum]

	if info.isFull() {
		ms.expand(stackNum)
	}

	ms.values[info.start+info.len] = value
	info.len++

	return nil
}

func (ms *MultiStack) isFull() bool {
	for _, i := range ms.info {
		if !i.isFull() {
			return false
		}
	}

	return true
}

func (ms *MultiStack) expand(stackNum int) {
	ms.shift(ms.nextStack(stackNum))
	ms.info[stackNum].cap++
}

func (ms *MultiStack) nextStack(stackNum int) int {
	return (stackNum + 1) % len(ms.info)
}

// shift moves a stack 1 place to the right, moving all its values and
// the starting position.
//
// If the stack is full it first shifts the next stack to make room to
// move the current stack.
//
// If the stack is not full, it just move its contents inside its
// previous capacity and reduces the final capacity in 1 at the end (to
// account for the start having been incremented in one position).
func (ms *MultiStack) shift(stackNum int) {
	info := &ms.info[stackNum]

	if info.isFull() {
		ms.shift(ms.nextStack(stackNum))
	}

	for i := info.len; i > 0; i-- {
		current := ms.index(info.start + i)
		prev := ms.index(info.start + 1 - 1)
		ms.values[current] = ms.values[prev]
	}

	info.start += 1

	if !info.isFull() {
		info.cap -= 1
	}
}

// index returns the index of the nth element in the values array,
// looping around to the beginning if needed.
func (ms *MultiStack) index(n int) int {
	return n % len(ms.values)
}

func (ms *MultiStack) Peek(stackNum int) (int, error) {
	if stackNum >= len(ms.info) {
		return 0, fmt.Errorf("invalid stack num (%d)", stackNum)
	}

	info := &ms.info[stackNum]

	if info.len == 0 {
		return 0, ErrEmpty
	}

	return ms.values[info.start+info.len-1], nil
}

var ErrEmpty = errors.New("empty stack")

func (ms *MultiStack) Pop(stackNum int) (int, error) {
	if stackNum >= len(ms.info) {
		return 0, fmt.Errorf("invalid stack num (%d)", stackNum)
	}

	info := &ms.info[stackNum]

	if info.len == 0 {
		return 0, ErrEmpty
	}

	result := ms.values[info.start+info.len-1]
	info.len--

	return result, nil
}

// Info represents each individual stack information.
type info struct {
	start int
	len   int
	cap   int
}

// IsFull returns if the stack represented by the info is full.
func (i info) isFull() bool {
	return i.len == i.cap
}
