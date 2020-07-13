package ch02_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci-6th/ch02"
)

func TestStringOfNewLinkedList(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input []int
		want  string
	}{
		{[]int(nil), ""},
		{[]int{}, ""},
		{[]int{0}, "[0]"},
		{[]int{0, 1}, "[0 1]"},
		{[]int{0, 1, 2}, "[0 1 2]"},
	}

	for _, test := range subtests {
		test := test
		desc := fmt.Sprintf("%d", test.input)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			l := ch02.NewLinkedList(test.input)
			got := l.String()

			if test.want != got {
				t.Fatalf("\nwant %q\n got %q", test.want, got)
			}
		})
	}
}

func TestClonedAreEqual(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input []int
	}{
		{[]int(nil)},
		{[]int{}},
		{[]int{0}},
		{[]int{0, 1}},
		{[]int{0, 1, 2}},
	}

	for _, test := range subtests {
		test := test
		desc := fmt.Sprintf("%q", test.input)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			l := ch02.NewLinkedList(test.input)
			cloned := l.Clone()

			if !cloned.Equals(l) {
				t.Fatalf("\noriginal %q\n cloned %q", l, cloned)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		a, b []int
		want bool
	}{
		{[]int(nil), []int(nil), true},
		{[]int{}, []int{}, true},
		{[]int(nil), []int{}, true},
		{[]int{}, []int(nil), true},

		{[]int{0}, []int{0}, true},
		{[]int{0}, []int{1}, false},
		{[]int{0}, []int{}, false},
		{[]int{0}, []int(nil), false},

		{[]int{0, 1}, []int{0, 1}, true},
		{[]int{0, 1}, []int{1, 0}, false},
		{[]int{0, 1}, []int{0, 0}, false},
		{[]int{0, 1}, []int{1, 1}, false},
		{[]int{0, 1}, []int{0}, false},
		{[]int{0, 1}, []int{1}, false},
		{[]int{0, 1}, []int{}, false},
		{[]int{0, 1}, []int(nil), false},
	}

	for _, test := range subtests {
		test := test
		desc := fmt.Sprintf("%q %q", test.a, test.b)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			la := ch02.NewLinkedList(test.a)
			lb := ch02.NewLinkedList(test.b)

			got := la.Equals(lb)
			if test.want != got {
				t.Fatalf("\na is %q\nb is %q", la, lb)
			}

			got = lb.Equals(la)
			if test.want != got {
				t.Fatalf("(reversed)\na is %q\nb is %q", la, lb)
			}
		})
	}
}
