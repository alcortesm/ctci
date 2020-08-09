package p04

import (
	"github.com/alcortesm/ctci/ch03"
)

// Queue implements a queue using two ch03.Stacks. Its zero value is
// safe to use.
type Queue struct {
	// newest inserted values, newest on top
	newest ch03.Stack
	// oldest inserted values, in reverse order, this is, oldest on top
	oldest ch03.Stack
}

func (q *Queue) Enqueue(v int) {
	q.newest.Push(v)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.oldest.Len() == 0 {
		q.shift()
	}

	return q.oldest.Pop()
}

func (q *Queue) Peek() (int, bool) {
	if q.oldest.Len() == 0 {
		q.shift()
	}

	return q.oldest.Pop()
}

// shift moves the newest contents into oldest; they end up in reverse
// order.
func (q *Queue) shift() {
	for {
		v, ok := q.newest.Pop()
		if !ok {
			break
		}

		q.oldest.Push(v)
	}
}

func (q *Queue) Len() int {
	return q.newest.Len() + q.oldest.Len()
}

// Dump returns the contents of the queue in arrival order.
func (q *Queue) Dump() []int {
	result := make([]int, 0, q.Len())

	// append oldest values, reversing their order
	for i := 0; i < q.oldest.Len(); i++ {
		v := (q.oldest)[q.oldest.Len()-1-i]
		result = append(result, v)
	}

	// then append newest values
	result = append(result, (q.newest)...)

	return result
}
