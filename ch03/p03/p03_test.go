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

	checkEmpty(t, s)
	checkPop(t, s, 0, false, [][]int{})

	checkPush(t, s, 0, [][]int{{0}})
	checkPush(t, s, 1, [][]int{{0, 1}})
	checkPush(t, s, 2, [][]int{{0, 1, 2}})

	checkPop(t, s, 2, true, [][]int{{0, 1}})
	checkPop(t, s, 1, true, [][]int{{0}})
	checkPop(t, s, 0, true, [][]int{})
	checkPop(t, s, 0, false, [][]int{})
	checkEmpty(t, s)

	checkPush(t, s, 0, [][]int{{0}})
	checkPush(t, s, 1, [][]int{{0, 1}})
	checkPush(t, s, 2, [][]int{{0, 1, 2}})
	checkPush(t, s, 3, [][]int{{0, 1, 2}, {3}})
	checkPush(t, s, 4, [][]int{{0, 1, 2}, {3, 4}})
	checkPush(t, s, 5, [][]int{{0, 1, 2}, {3, 4, 5}})
	checkPush(t, s, 6, [][]int{{0, 1, 2}, {3, 4, 5}, {6}})
	checkPush(t, s, 7, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7}})

	checkPop(t, s, 7, true, [][]int{{0, 1, 2}, {3, 4, 5}, {6}})
	checkPop(t, s, 6, true, [][]int{{0, 1, 2}, {3, 4, 5}})
	checkPop(t, s, 5, true, [][]int{{0, 1, 2}, {3, 4}})
	checkPop(t, s, 4, true, [][]int{{0, 1, 2}, {3}})
	checkPop(t, s, 3, true, [][]int{{0, 1, 2}})
	checkPop(t, s, 2, true, [][]int{{0, 1}})
	checkPop(t, s, 1, true, [][]int{{0}})
	checkPop(t, s, 0, true, [][]int{})

	checkPop(t, s, 0, false, [][]int{})
	checkEmpty(t, s)
}

func checkPush(t *testing.T, s *p03.SetOfStacks, v int, dump [][]int) {
	t.Helper()

	s.Push(v)

	got := s.Dump()
	if diff := cmp.Diff(dump, got); diff != "" {
		t.Errorf("(-want +got)\n%s", diff)
	}
}

func checkPop(t *testing.T, s *p03.SetOfStacks, v int, ok bool, dump [][]int) {
	t.Helper()

	gotValue, gotOK := s.Pop()
	if ok != gotOK {
		t.Errorf("wron OK: want %t got %t", ok, gotOK)
	}

	if v != gotValue {
		t.Errorf("wrong value: want %d, got %d", v, gotValue)
	}

	got := s.Dump()
	if diff := cmp.Diff(dump, got); diff != "" {
		t.Errorf("(-want +got)\n%s", diff)
	}
}

func checkEmpty(t *testing.T, s *p03.SetOfStacks) {
	t.Helper()

	if got := s.Len(); 0 != got {
		t.Errorf("not empty, len %d", got)
	}
}

func TestSetOfStacks_PopAt(t *testing.T) {
	s, err := p03.NewSetOfStacks(3)
	if err != nil {
		t.Fatal(err)
	}

	// empty
	checkEmpty(t, s)
	checkPopAt(t, s, 0, 0, false, [][]int{})
	checkPopAt(t, s, 1, 0, false, [][]int{})

	// popAt 0 from [0]
	checkPush(t, s, 0, [][]int{{0}})
	checkPopAt(t, s, 0, 0, true, [][]int{})
	checkEmpty(t, s)
	checkPopAt(t, s, 0, 0, false, [][]int{})
	checkPopAt(t, s, 1, 0, false, [][]int{})

	// popAt 1 from [0]
	checkPush(t, s, 0, [][]int{{0}})
	checkPopAt(t, s, 1, 0, false, [][]int{{0}})
	checkPopAt(t, s, 0, 0, true, [][]int{})
	checkEmpty(t, s)
	checkPopAt(t, s, 0, 0, false, [][]int{})
	checkPopAt(t, s, 1, 0, false, [][]int{})

	// popAt 0 from [0, 1, 2]
	checkPush(t, s, 0, [][]int{{0}})
	checkPush(t, s, 1, [][]int{{0, 1}})
	checkPush(t, s, 2, [][]int{{0, 1, 2}})

	checkPopAt(t, s, 1, 0, false, [][]int{{0, 1, 2}})
	checkPopAt(t, s, 0, 2, true, [][]int{{0, 1}})

	checkPopAt(t, s, 1, 0, false, [][]int{{0, 1}})
	checkPopAt(t, s, 0, 1, true, [][]int{{0}})

	checkPopAt(t, s, 1, 0, false, [][]int{{0}})
	checkPopAt(t, s, 0, 0, true, [][]int{})

	checkEmpty(t, s)
	checkPopAt(t, s, 0, 0, false, [][]int{})
	checkPopAt(t, s, 1, 0, false, [][]int{})

	// popAt 0 from [0, 1, 2], [3, 4, 5], [6, 7]
	checkPush(t, s, 0, [][]int{{0}})
	checkPush(t, s, 1, [][]int{{0, 1}})
	checkPush(t, s, 2, [][]int{{0, 1, 2}})
	checkPush(t, s, 3, [][]int{{0, 1, 2}, {3}})
	checkPush(t, s, 4, [][]int{{0, 1, 2}, {3, 4}})
	checkPush(t, s, 5, [][]int{{0, 1, 2}, {3, 4, 5}})
	checkPush(t, s, 6, [][]int{{0, 1, 2}, {3, 4, 5}, {6}})
	checkPush(t, s, 7, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7}})

	checkPopAt(t, s, 0, 2, true, [][]int{{0, 1, 3}, {4, 5, 6}, {7}})
	checkPopAt(t, s, 0, 3, true, [][]int{{0, 1, 4}, {5, 6, 7}})
	checkPopAt(t, s, 0, 4, true, [][]int{{0, 1, 5}, {6, 7}})
	checkPopAt(t, s, 0, 5, true, [][]int{{0, 1, 6}, {7}})
	checkPopAt(t, s, 0, 6, true, [][]int{{0, 1, 7}})

	// different popAt from [0, 1, 7], [8, 9, 10], [11]
	checkPush(t, s, 8, [][]int{{0, 1, 7}, {8}})
	checkPush(t, s, 9, [][]int{{0, 1, 7}, {8, 9}})
	checkPush(t, s, 10, [][]int{{0, 1, 7}, {8, 9, 10}})
	checkPush(t, s, 11, [][]int{{0, 1, 7}, {8, 9, 10}, {11}})

	checkPopAt(t, s, 1, 10, true, [][]int{{0, 1, 7}, {8, 9, 11}})
	checkPopAt(t, s, 0, 7, true, [][]int{{0, 1, 8}, {9, 11}})
	checkPopAt(t, s, 1, 11, true, [][]int{{0, 1, 8}, {9}})
	checkPopAt(t, s, 0, 8, true, [][]int{{0, 1, 9}})
	checkPopAt(t, s, 1, 0, false, [][]int{{0, 1, 9}})
	checkPopAt(t, s, 0, 9, true, [][]int{{0, 1}})
	checkPop(t, s, 1, true, [][]int{{0}})
	checkPop(t, s, 0, true, [][]int{})

	checkEmpty(t, s)
	checkPopAt(t, s, 0, 0, false, [][]int{})
	checkPopAt(t, s, 1, 0, false, [][]int{})
}

func checkPopAt(t *testing.T, s *p03.SetOfStacks, stackNum int, v int, ok bool, dump [][]int) {
	t.Helper()

	gotValue, gotOK := s.PopAt(stackNum)
	if ok != gotOK {
		t.Errorf("wron OK: want %t got %t", ok, gotOK)
	}

	if v != gotValue {
		t.Errorf("wrong value: want %d, got %d", v, gotValue)
	}

	got := s.Dump()
	if diff := cmp.Diff(dump, got); diff != "" {
		t.Errorf("(-want +got)\n%s", diff)
	}
}
