package internal

import (
	"errors"
)

var ErrEmpty = errors.New("empty")

type Intake struct {
	Item           interface{}
	AdmissionOrder int // a monotonic counter
}

func (i *Intake) IsOldest(j *Intake) bool {
	return i.AdmissionOrder < j.AdmissionOrder
}

type Queue struct {
	first, last *node
}

type node struct {
	intake *Intake
	next   *node
}

func (q *Queue) Enqueue(i *Intake) {
	n := &node{intake: i}

	if q.first == nil {
		q.first = n
		q.last = n
		return
	}

	q.last.next = n
	q.last = n
}

func (q *Queue) Dequeue() (*Intake, error) {
	if q.first == nil {
		return nil, ErrEmpty
	}

	tmp := q.first
	q.first = q.first.next
	tmp.next = nil // helps garbage collector

	return tmp.intake, nil
}

func (q *Queue) Peek() (*Intake, error) {
	if q.first == nil {
		return nil, ErrEmpty
	}

	return q.first.intake, nil
}
