package ch03_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci/ch03"
	"github.com/google/go-cmp/cmp"
)

func TestMultiStack(t *testing.T) {
	t.Parallel()

	subtests := map[string]func(t *testing.T){
		"new stacks are empty":        msNewAreEmpty,
		"can add to empty stack":      msPushToEmpty,
		"one stack can use all space": msOneCanUseAllSpace,
		"can fill all":                msCanFillAll,
		"expand shifts next":          msExpandShiftsOne,
		"expand shifts next circular": msExpandShiftsOneCircular,
		"expand shifts all":           msExpandShiftsAll,
		"expand shifts all circular":  msExpandShiftsAllCircular,
		"can peek":                    msCanPeek,
		"can pop":                     msCanPop,
	}

	for name, fn := range subtests {
		fn := fn

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			fn(t)
		})
	}
}

func msNewAreEmpty(t *testing.T) {
	subtests := []struct {
		numStacks   int
		defaultSize int
		want        [][]int
	}{
		{
			numStacks:   3,
			defaultSize: 1,
			want:        [][]int{{}, {}, {}},
		},
		{
			numStacks:   3,
			defaultSize: 2,
			want:        [][]int{{}, {}, {}},
		},
		{
			numStacks:   4,
			defaultSize: 2,
			want:        [][]int{{}, {}, {}, {}},
		},
	}

	for _, test := range subtests {
		test := test

		name := fmt.Sprintf("%d %d", test.numStacks, test.defaultSize)

		t.Run(name, func(t *testing.T) {
			ms := mustNewMultiStack(t, test.numStacks, test.defaultSize)
			check(t, ms, test.want)
		})
	}
}

func mustNewMultiStack(t *testing.T, n, d int) *ch03.MultiStack {
	t.Helper()

	result, err := ch03.NewMultiStack(n, d)
	if err != nil {
		t.Fatal(err)
	}

	return result
}

func check(t *testing.T, ms *ch03.MultiStack, want [][]int) {
	t.Helper()

	got := ms.PeekAll()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("(-want +got):\n%s", diff)
	}
}

func msPushToEmpty(t *testing.T) {
	subtests := []struct {
		ms       *ch03.MultiStack
		stackNum int
		value    int
		want     [][]int
	}{
		{
			ms:       mustNewMultiStack(t, 3, 1),
			stackNum: 0,
			value:    1,
			want:     [][]int{{1}, {}, {}},
		},
		{
			ms:       mustNewMultiStack(t, 3, 1),
			stackNum: 1,
			value:    1,
			want:     [][]int{{}, {1}, {}},
		},
		{
			ms:       mustNewMultiStack(t, 3, 1),
			stackNum: 2,
			value:    1,
			want:     [][]int{{}, {}, {1}},
		},
		{
			ms:       mustNewMultiStack(t, 3, 2),
			stackNum: 0,
			value:    1,
			want:     [][]int{{1}, {}, {}},
		},
		{
			ms:       mustNewMultiStack(t, 3, 2),
			stackNum: 1,
			value:    1,
			want:     [][]int{{}, {1}, {}},
		},
		{
			ms:       mustNewMultiStack(t, 3, 2),
			stackNum: 2,
			value:    1,
			want:     [][]int{{}, {}, {1}},
		},
	}

	for _, test := range subtests {
		test := test

		name := fmt.Sprintf("%s %d %d", test.ms, test.stackNum, test.value)

		t.Run(name, func(t *testing.T) {
			test.ms.Push(test.stackNum, test.value)
			check(t, test.ms, test.want)
		})
	}
}

func msOneCanUseAllSpace(t *testing.T) {
	ms := mustNewMultiStack(t, 3, 2)
	ms.Push(0, 1)
	ms.Push(0, 2)
	ms.Push(0, 3)
	ms.Push(0, 4)
	ms.Push(0, 5)
	ms.Push(0, 6)
	want := [][]int{{1, 2, 3, 4, 5, 6}, {}, {}}
	check(t, ms, want)
}

func msCanFillAll(t *testing.T) {
	ms := mustNewMultiStack(t, 3, 2)
	ms.Push(0, 1)
	ms.Push(0, 2)
	ms.Push(1, 3)
	ms.Push(1, 4)
	ms.Push(2, 5)
	ms.Push(2, 6)
	want := [][]int{{1, 2}, {3, 4}, {5, 6}}
	check(t, ms, want)
}

func msExpandShiftsOne(t *testing.T) {
	ms := mustNewMultiStack(t, 3, 2)
	ms.Push(0, 1)
	ms.Push(0, 2)
	ms.Push(1, 3)
	ms.Push(1, 4)
	ms.Push(2, 5)
	ms.Push(2, 6)
	want := [][]int{{1, 2}, {3, 4}, {5, 6}}
	check(t, ms, want)
}

func msExpandShiftsOneCircular(t *testing.T) {
	ms := mustNewMultiStack(t, 3, 2)
	ms.Push(0, 0)
	ms.Push(1, 2)
	ms.Push(1, 3)
	ms.Push(2, 4)
	ms.Push(1, 44)
	ms.Push(1, 55)
	want := [][]int{{0}, {2, 3, 44, 55}, {4}}
	check(t, ms, want)
}

func msExpandShiftsAll(t *testing.T) {
	ms := mustNewMultiStack(t, 3, 2)
	ms.Push(0, 1)
	ms.Push(0, 2)
	ms.Push(1, 3)
	ms.Push(2, 5)
	ms.Push(0, 33)
	ms.Push(0, 44)
	want := [][]int{{1, 2, 33, 44}, {3}, {5}}
	check(t, ms, want)
}

func msExpandShiftsAllCircular(t *testing.T) {
	ms := mustNewMultiStack(t, 4, 2)
	ms.Push(1, 2)
	ms.Push(1, 3)
	ms.Push(2, 4)
	ms.Push(3, 6)
	want := [][]int{{}, {2, 3}, {4}, {6}}
	check(t, ms, want)

	ms.Push(1, 44)
	want = [][]int{{}, {2, 3, 44}, {4}, {6}}
	check(t, ms, want)

	ms.Push(1, 55)
	want = [][]int{{}, {2, 3, 44, 55}, {4}, {6}}
	check(t, ms, want)

	ms.Push(1, 66)
	want = [][]int{{}, {2, 3, 44, 55, 66}, {4}, {6}}
	check(t, ms, want)

	ms.Push(1, 77)
	want = [][]int{{}, {2, 3, 44, 55, 66, 77}, {4}, {6}}
	check(t, ms, want)
}

func msCanPeek(t *testing.T) {
	ms := mustNewMultiStack(t, 3, 2)
	ms.Push(0, 0)
	ms.Push(0, 1)
	ms.Push(1, 2)
	ms.Push(2, 4)

	want := 1
	got, err := ms.Peek(0)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

	want = 2
	got, err = ms.Peek(1)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

	want = 4
	got, err = ms.Peek(2)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func msCanPop(t *testing.T) {
	ms := mustNewMultiStack(t, 3, 2)
	ms.Push(0, 0)
	ms.Push(0, 1)
	ms.Push(1, 2)
	ms.Push(2, 4)

	want := 1
	got, err := ms.Pop(0)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

	want = 0
	got, err = ms.Pop(0)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

	want = 2
	got, err = ms.Pop(1)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

	want = 4
	got, err = ms.Pop(2)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
