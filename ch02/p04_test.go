package ch02_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci/ch02"
)

func TestPartition(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		list []int
		x    int
	}{
		{list: []int{0}, x: 0},

		{list: []int{0, 1}, x: 0},
		{list: []int{0, 1}, x: 1},
		{list: []int{0, 1}, x: 2},

		{list: []int{1, 0}, x: 0},
		{list: []int{1, 0}, x: 1},
		{list: []int{1, 0}, x: 2},

		{list: []int{0, 1, 2}, x: 0},
		{list: []int{0, 1, 2}, x: 1},
		{list: []int{0, 1, 2}, x: 2},
		{list: []int{0, 1, 2}, x: 3},

		{list: []int{2, 1, 0}, x: 0},
		{list: []int{2, 1, 0}, x: 1},
		{list: []int{2, 1, 0}, x: 2},
		{list: []int{2, 1, 0}, x: 3},

		{list: []int{0, 1, 2, 3}, x: 0},
		{list: []int{0, 1, 2, 3}, x: 1},
		{list: []int{0, 1, 2, 3}, x: 2},
		{list: []int{0, 1, 2, 3}, x: 3},
		{list: []int{0, 1, 2, 3}, x: 4},

		{list: []int{4, 3, 3, 1}, x: 0},
		{list: []int{4, 3, 3, 1}, x: 1},
		{list: []int{4, 3, 3, 1}, x: 2},
		{list: []int{4, 3, 3, 1}, x: 3},
		{list: []int{4, 3, 3, 1}, x: 4},

		{list: []int{3, 5, 8, 5, 10, 2, 1}, x: 5},
	}

	for _, test := range subtests {
		test := test

		l := ch02.NewLinkedList(test.list)
		desc := fmt.Sprintf("%s %d", l, test.x)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			c := l.Clone()
			ch02.Partition(&c, test.x)

			if !partitioned(c, test.x) {
				t.Errorf("not partitioned: %q", l)
			}

		})
	}
}

func partitioned(l ch02.LinkedList, x int) bool {

	// find the partition index
	partitionIndex := -1
	{
		index := 0
		for n := l.First; n != nil; n = n.Next {
			if n.Data >= x {
				partitionIndex = index
			}

			index++
		}
	}

	// no data in the list is bigger or equal to x,
	// so the list is always partitioned at x no matter the order of its
	// elements
	if partitionIndex == -1 {
		return true
	}

	// after that index all data must be greater or equal to x
	index := 0
	for n := l.First; n != nil; n = n.Next {
		if index >= partitionIndex {
			if n.Data < x {
				return false
			}
		}

		index++
	}

	return true
}
