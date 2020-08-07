package ch02_test

import (
	"testing"

	"github.com/alcortesm/ctci/ch02"
)

func TestRemoveDups(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input []int
		want  []int
	}{
		{input: []int{}, want: []int{}},
		{input: []int{0}, want: []int{0}},
		{input: []int{0, 1}, want: []int{0, 1}},

		{input: []int{0, 0}, want: []int{0}},
		{input: []int{0, 0, 0}, want: []int{0}},

		{input: []int{0, 1, 0}, want: []int{0, 1}},
		{input: []int{0, 1, 0, 2}, want: []int{0, 1, 2}},
		{input: []int{0, 1, 0, 2, 2, 0, 0, 2, 3, 1}, want: []int{0, 1, 2, 3}},
	}

	for _, test := range subtests {
		test := test

		l := ch02.NewLinkedList(test.input)
		want := ch02.NewLinkedList(test.want)

		t.Run(l.String(), func(t *testing.T) {
			t.Parallel()

			c := l.Clone()
			ch02.RemoveDupsUsingSpace(c)
			if !c.Equals(want) {
				t.Errorf("(using space)\nwant %q\n got %q", want, c)
			}

			cc := l.Clone()
			ch02.RemoveDupsUsingTime(cc)
			if !cc.Equals(want) {
				t.Errorf("(using time)\nwant %q\n got %q", want, cc)
			}

		})
	}
}
