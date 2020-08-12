package p01_test

import (
	"testing"

	"github.com/alcortesm/ctci/ch04/p01"
)

func TestQueue(t *testing.T) {
	type person struct {
		name string
	}

	var v0 int = 42
	var v1 string = "foo"
	var v2 *person = &person{name: "a"}

	q := new(p01.Queue)
	checkIsEmpty(t, q, true)

	q.Enqueue(v0)
	checkIsEmpty(t, q, false)

	checkDequeue(t, q, v0)
	checkIsEmpty(t, q, true)

	q.Enqueue(v0)
	q.Enqueue(v1)
	q.Enqueue(v2)
	checkIsEmpty(t, q, false)

	checkDequeue(t, q, v0)
	checkIsEmpty(t, q, false)

	checkDequeue(t, q, v1)
	checkIsEmpty(t, q, false)

	checkDequeue(t, q, v2)
	checkIsEmpty(t, q, true)
}

func checkIsEmpty(t *testing.T, q *p01.Queue, want bool) {
	t.Helper()

	if got := q.IsEmpty(); got != want {
		t.Fatalf("IsEmpty check: want %t, got %t", want, got)
	}
}

func checkDequeue(t *testing.T, q *p01.Queue, want interface{}) {
	t.Helper()

	if got := q.Dequeue(); got != want {
		t.Fatalf("Dequeue check:\nwant %#v\n got %#v", want, got)
	}
}
