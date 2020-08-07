package p02_test

import (
	"testing"

	"github.com/alcortesm/ctci/ch03/p02"
)

func TestStackWithMin(t *testing.T) {
	t.Parallel()

	s := p02.NewStackWithMin()
	checkEmpty(t, s)

	s.Push(2)
	checkPeekAndMin(t, s, 2, 2)

	s.Push(4)
	checkPeekAndMin(t, s, 4, 2)

	s.Push(1)
	checkPeekAndMin(t, s, 1, 1)

	checkPop(t, s, 1)
	checkPeekAndMin(t, s, 4, 2)

	checkPop(t, s, 4)
	checkPeekAndMin(t, s, 2, 2)

	checkPop(t, s, 2)
	checkEmpty(t, s)

	s.Push(1)
	checkPeekAndMin(t, s, 1, 1)

	s.Push(2)
	checkPeekAndMin(t, s, 2, 1)

	s.Push(1)
	checkPeekAndMin(t, s, 1, 1)

	s.Push(1)
	checkPeekAndMin(t, s, 1, 1)

	checkPop(t, s, 1)
	checkPeekAndMin(t, s, 1, 1)

	checkPop(t, s, 1)
	checkPeekAndMin(t, s, 2, 1)

	checkPop(t, s, 2)
	checkPeekAndMin(t, s, 1, 1)

	checkPop(t, s, 1)
	checkEmpty(t, s)
}

func checkPeekAndMin(t *testing.T, s *p02.StackWithMin, top, min int) {
	t.Helper()

	got, ok := s.Peek()
	if !ok {
		t.Fatal("empty")
	}

	if got != top {
		t.Errorf("wrong top: want %d, got %d", top, got)
	}

	got, ok = s.Min()
	if !ok {
		t.Fatal("empty")
	}

	if got != min {
		t.Errorf("wrong min: want %d, got %d", min, got)
	}
}

func checkPop(t *testing.T, s *p02.StackWithMin, want int) {
	t.Helper()

	got, ok := s.Pop()
	if !ok {
		t.Fatal("empty")
	}

	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func checkEmpty(t *testing.T, s *p02.StackWithMin) {
	t.Helper()

	if got, ok := s.Peek(); ok {
		t.Fatalf("want empty, but top is %d", got)
	}

	if got, ok := s.Min(); ok {
		t.Fatalf("want empty, but min is %d", got)
	}
}
