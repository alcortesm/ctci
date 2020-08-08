package p03_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/alcortesm/ctci/ch03/p03"
)

func TestSetOfStacks(t *testing.T) {
	s, err := p03.NewSetOfStacks(3)
	if err != nil {
		t.Fatal(err)
	}

	checkDump(t, s, [][]int{})

	s.Push(0)
	checkDump(t, s, [][]int{{0}})
	s.Push(1)
	checkDump(t, s, [][]int{{0, 1}})
	s.Push(2)
	checkDump(t, s, [][]int{{0, 1, 2}})

	checkPop(t, s, 2)
	checkDump(t, s, [][]int{{0, 1}})
	checkPop(t, s, 1)
	checkDump(t, s, [][]int{{0}})
	checkPop(t, s, 0)
	checkEmpty(t, s)

	s.Push(0)
	s.Push(1)
	s.Push(2)
	checkDump(t, s, [][]int{{0, 1, 2}})
	s.Push(3)
	checkDump(t, s, [][]int{{0, 1, 2}, {3}})
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	checkDump(t, s, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7}})

	checkPop(t, s, 7)
	checkPop(t, s, 6)
	checkDump(t, s, [][]int{{0, 1, 2}, {3, 4, 5}})
	checkPop(t, s, 5)
	checkPop(t, s, 4)
	checkDump(t, s, [][]int{{0, 1, 2}, {3}})
	checkPop(t, s, 3)
	checkPop(t, s, 2)
	checkDump(t, s, [][]int{{0, 1}})
	checkPop(t, s, 1)
	checkPop(t, s, 0)
	checkDump(t, s, [][]int{})
	checkEmpty(t, s)
}

func checkDump(t *testing.T, s *p03.SetOfStacks, want [][]int) {
	t.Helper()

	got := s.Dump()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("(-want +got)\n%s", diff)
	}
}

func checkPop(t *testing.T, s *p03.SetOfStacks, want int) {
	t.Helper()

	got, ok := s.Pop()
	if !ok {
		t.Fatal("empty")
	}

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func checkEmpty(t *testing.T, s *p03.SetOfStacks) {
	t.Helper()

	if got := s.Len(); 0 != got {
		t.Errorf("not empty, len %d", got)
	}
}
