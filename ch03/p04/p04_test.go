package p04_test

import (
	"testing"

	"github.com/alcortesm/ctci/ch03/p04"
	"github.com/google/go-cmp/cmp"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	q := &p04.Queue{}

	// empty
	checkDump(t, q, []int{})
	checkLen(t, q, 0)

	// dequeue from empty
	//checkDequeue(t, q, 0, false, []int{})
	checkLen(t, q, 0)

	// enqueue 0, then dequeue
	checkEnqueue(t, q, 0, []int{0})
	checkLen(t, q, 1)
	checkDequeue(t, q, 0, true, []int{})
	checkLen(t, q, 0)

	// enqueue 0, 1, then dequeue both
	checkEnqueue(t, q, 0, []int{0})
	checkEnqueue(t, q, 1, []int{0, 1})
	checkLen(t, q, 2)
	checkDequeue(t, q, 0, true, []int{1})
	checkLen(t, q, 1)
	checkDequeue(t, q, 1, true, []int{})
	checkLen(t, q, 0)

	// enqueue 0, 1, 2, then dequeue and enqueue several times
	checkEnqueue(t, q, 0, []int{0})
	checkLen(t, q, 1)
	checkEnqueue(t, q, 1, []int{0, 1})
	checkLen(t, q, 2)
	checkEnqueue(t, q, 2, []int{0, 1, 2})
	checkLen(t, q, 3)
	checkDequeue(t, q, 0, true, []int{1, 2})
	checkLen(t, q, 2)
	checkEnqueue(t, q, 3, []int{1, 2, 3})
	checkLen(t, q, 3)
	checkEnqueue(t, q, 4, []int{1, 2, 3, 4})
	checkLen(t, q, 4)
	checkDequeue(t, q, 1, true, []int{2, 3, 4})
	checkLen(t, q, 3)
	checkDequeue(t, q, 2, true, []int{3, 4})
	checkLen(t, q, 2)
	checkDequeue(t, q, 3, true, []int{4})
	checkLen(t, q, 1)
	checkEnqueue(t, q, 5, []int{4, 5})
	checkLen(t, q, 2)
	checkDequeue(t, q, 4, true, []int{5})
	checkLen(t, q, 1)
	checkDequeue(t, q, 5, true, []int{})
	checkLen(t, q, 0)

	// enqueue and dequeue many elements
	n := 10_000
	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < n; i++ {
		got, ok := q.Dequeue()
		if !ok {
			t.Fatalf("no ok at i=%d", i)
		}

		if got != i {
			t.Errorf("want %d, got %d", i, got)
		}
	}
}

func checkDump(t *testing.T, q *p04.Queue, want []int) {
	t.Helper()

	got := q.Dump()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("(-want +got)\n%s", diff)
	}
}

func checkLen(t *testing.T, q *p04.Queue, want int) {
	t.Helper()

	got := q.Len()
	if got != want {
		t.Errorf("wrong len: want %d, got %d", want, got)
	}
}

func checkEnqueue(t *testing.T, q *p04.Queue, v int, dump []int) {
	t.Helper()

	q.Enqueue(v)
	checkDump(t, q, dump)
}

func checkDequeue(t *testing.T, q *p04.Queue, v int, ok bool, dump []int) {
	t.Helper()

	gotValue, gotOK := q.Dequeue()
	if ok != gotOK {
		t.Errorf("wron OK: want %t got %t", ok, gotOK)
	}

	if v != gotValue {
		t.Errorf("wrong value: want %d, got %d", v, gotValue)
	}

	checkDump(t, q, dump)
}
