package ch03_test

import (
	"testing"

	"github.com/alcortesm/ctci/ch03"
	"github.com/google/go-cmp/cmp"
)

func TestLinkedStack(t *testing.T) {
	s, err := ch03.NewLinkedStack(3)
	if err != nil {
		t.Fatal(err)
	}

	// empty stack
	checkPop(t, s, 0, false, []int{})
	checkPeek(t, s, 0, false, []int{})
	checkRemoveBottom(t, s, 0, false, []int{})

	// push 0
	checkPush(t, s, 0, true, []int{0})
	checkPeek(t, s, 0, true, []int{0})

	// empty by poping
	checkPop(t, s, 0, true, []int{})
	checkPop(t, s, 0, false, []int{})
	checkPeek(t, s, 0, false, []int{})
	checkRemoveBottom(t, s, 0, false, []int{})

	// push 0 again
	checkPush(t, s, 0, true, []int{0})
	checkPeek(t, s, 0, true, []int{0})

	// empty by removing the bottom
	checkRemoveBottom(t, s, 0, true, []int{})
	checkPeek(t, s, 0, false, []int{})
	checkPop(t, s, 0, false, []int{})
	checkRemoveBottom(t, s, 0, false, []int{})

	// push until full: 0, 1, 2
	checkPush(t, s, 0, true, []int{0})
	checkPush(t, s, 1, true, []int{0, 1})
	checkPush(t, s, 2, true, []int{0, 1, 2})
	checkPush(t, s, 3, false, []int{0, 1, 2})

	// empty by poping and removing from bottom
	checkRemoveBottom(t, s, 0, true, []int{1, 2})
	checkPeek(t, s, 2, true, []int{1, 2})

	checkPop(t, s, 2, true, []int{1})
	checkPeek(t, s, 1, true, []int{1})

	checkPush(t, s, 100, true, []int{1, 100})
	checkPeek(t, s, 100, true, []int{1, 100})

	checkRemoveBottom(t, s, 1, true, []int{100})
	checkPeek(t, s, 100, true, []int{100})

	checkPop(t, s, 100, true, []int{})
	checkPop(t, s, 0, false, []int{})
	checkPeek(t, s, 0, false, []int{})
	checkRemoveBottom(t, s, 0, false, []int{})
}

func checkPush(t *testing.T, s *ch03.LinkedStack, v int, want bool, dump []int) {
	t.Helper()

	if got := s.Push(v); got != want {
		t.Errorf("wrong push return value: want %t, got %t", want, got)
	}

	got := s.Dump()
	if diff := cmp.Diff(dump, got); diff != "" {
		t.Errorf("wrong dump (-want +got)\n%s", diff)
	}
}

func checkPop(t *testing.T, s *ch03.LinkedStack, v int, ok bool, dump []int) {
	t.Helper()

	gotValue, gotOK := s.Pop()

	if gotOK != ok {
		t.Errorf("wrong pop return ok: want %t, got %t", ok, gotOK)
	}

	if gotValue != v {
		t.Errorf("wrong pop return value: want %d, got %d", v, gotValue)
	}

	got := s.Dump()
	if diff := cmp.Diff(dump, got); diff != "" {
		t.Errorf("wrong dump (-want +got)\n%s", diff)
	}
}

func checkPeek(t *testing.T, s *ch03.LinkedStack, v int, ok bool, dump []int) {
	t.Helper()

	gotValue, gotOK := s.Peek()

	if gotOK != ok {
		t.Errorf("wrong pop return ok: want %t, got %t", ok, gotOK)
	}

	if gotValue != v {
		t.Errorf("wrong pop return value: want %d, got %d", v, gotValue)
	}

	got := s.Dump()
	if diff := cmp.Diff(dump, got); diff != "" {
		t.Errorf("wrong dump (-want +got)\n%s", diff)
	}
}

func checkRemoveBottom(t *testing.T, s *ch03.LinkedStack, v int, ok bool, dump []int) {
	t.Helper()

	gotValue, gotOK := s.RemoveBottom()

	if gotOK != ok {
		t.Errorf("wrong pop return ok: want %t, got %t", ok, gotOK)
	}

	if gotValue != v {
		t.Errorf("wrong pop return value: want %d, got %d", v, gotValue)
	}

	got := s.Dump()
	if diff := cmp.Diff(dump, got); diff != "" {
		t.Errorf("wrong dump (-want +got)\n%s", diff)
	}
}
