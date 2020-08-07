package ch02_test

import (
	"testing"

	"github.com/alcortesm/ctci/ch02"
)

func TestPalindrome(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input []int
		want  bool
	}{
		{input: []int{}, want: true},
		{input: []int{0}, want: true},
		{input: []int{0, 0}, want: true},
		{input: []int{0, 1}, want: false},
		{input: []int{0, 0, 0}, want: true},
		{input: []int{0, 1, 0}, want: true},
		{input: []int{1, 1, 0}, want: false},
		{input: []int{0, 1, 2, 2, 1, 0}, want: true},
		{input: []int{0, 1, 2, 7, 1, 0}, want: false},
		{input: []int{0, 1, 2, 3, 2, 1, 0}, want: true},
		{input: []int{0, 1, 2, 3, 7, 1, 0}, want: false},
		{input: []int{0, 1, 2, 3, 2, 7, 0}, want: false},
		{input: []int{0, 1, 2, 3, 2, 1, 7}, want: false},
	}

	for _, test := range subtests {
		test := test

		l := ch02.NewLinkedList(test.input)

		t.Run(l.String(), func(t *testing.T) {
			t.Parallel()

			got := ch02.IsPalindrome(l)
			if got != test.want {
				t.Errorf("want %t, got %t", test.want, got)
			}
		})
	}
}
