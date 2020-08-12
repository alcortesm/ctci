package p01

type Queue struct {
	first, last *node
}

// A queue node.
type node struct {
	value interface{}
	next  *node
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Enqueue(v interface{}) {
	n := &node{value: v}

	if q.IsEmpty() {
		q.first = n
		q.last = n
		return
	}

	q.last.next = n
	q.last = n
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	result := q.first

	if result == q.last {
		q.last = nil
	}

	q.first = q.first.next
	result.next = nil

	return result.value
}

func (q *Queue) Dump() []interface{} {
	result := []interface{}{}

	for n := q.first; n != nil; n = n.next {
		result = append(result, n.value)
	}

	return result
}
