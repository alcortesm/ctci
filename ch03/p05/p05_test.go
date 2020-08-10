package p05_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/alcortesm/ctci/ch03/p05"
	"github.com/google/go-cmp/cmp"
)

func TestSort(t *testing.T) {
	t.Parallel()

	subtests := [][]int{
		// no need to sort
		nil,
		{},
		{0},
		// already sorted
		{1, 0},
		{2, 1, 0},
		{3, 2, 1, 0},
		// need to be sorted
		{0, 1},
		{0, 1, 2},
		{0, 1, 2, 3},
		// mix
		{0, 2, 1},
		{7, 3, 4, 0, 6, 1, 5, 2},
	}

	for _, input := range subtests {
		input := input
		name := fmt.Sprintf("%d", input)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// populate a stack with the input and,
			// call the Sort method and get a dump of the values.
			s := p05.Stack{}
			for _, v := range input {
				s.Push(v)
			}
			s.Sort()
			got := s.Dump()

			// Manually calculate the sorted version of the input.
			want := make([]int, len(input))
			copy(want, input)
			smallestFirst := func(i, j int) bool {
				return want[i] > want[j]
			}
			sort.Slice(want, smallestFirst)

			// compare Stack.Sort with our manually sorted data.
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("(-want +got)\n%s", diff)
			}
		})
	}
}
