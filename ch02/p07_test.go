package ch02_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci-6th/ch02"
)

func TestIntersect(t *testing.T) {
	t.Parallel()

	n := make([]ch02.Node, 9)
	for i := range n {
		n[i].Data = i
	}

	n[0].Next = &n[1]
	n[1].Next = &n[2]
	n[2].Next = &n[3]

	n[4].Next = &n[5]
	n[5].Next = &n[6]
	n[6].Next = &n[2]

	n[7].Next = &n[8]

	l1 := ch02.LinkedList{First: &n[0]}
	l2 := ch02.LinkedList{First: &n[4]}
	l3 := ch02.LinkedList{First: &n[7]}

	subtests := []struct {
		a, b ch02.LinkedList
		want *ch02.Node
	}{
		{a: l1, b: l1, want: &n[0]},
		{a: l2, b: l2, want: &n[4]},
		{a: l3, b: l3, want: &n[7]},

		{a: l1, b: l2, want: &n[2]},
		{a: l2, b: l1, want: &n[2]},

		{a: l1, b: l3, want: nil},
		{a: l3, b: l1, want: nil},

		{a: l2, b: l3, want: nil},
		{a: l3, b: l2, want: nil},
	}

	for _, test := range subtests {
		test := test
		desc := fmt.Sprintf("%s %s", test.a, test.b)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			got := ch02.Intersect(test.a, test.b)
			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
