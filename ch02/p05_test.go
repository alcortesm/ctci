package ch02_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci-6th/ch02"
)

func TestSum(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		a, b []int
		want []int
	}{
		{a: []int{}, b: []int{}, want: []int{0}},
		{a: []int{0}, b: []int{0}, want: []int{0}},
		{a: []int{1}, b: []int{2}, want: []int{3}},
		{a: []int{2, 7}, b: []int{3, 4, 1}, want: []int{5, 1, 2}},
		{a: []int{7, 1, 6}, b: []int{5, 9, 2}, want: []int{2, 1, 9}},
	}

	for _, test := range subtests {
		test := test

		a := ch02.NewLinkedList(test.a)
		b := ch02.NewLinkedList(test.b)
		desc := fmt.Sprintf("%s %s", a, b)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			got := ch02.Sum(a, b)
			want := ch02.NewLinkedList(test.want)

			if got.String() != want.String() {
				t.Errorf("\nwant: %s\n got: %s", want, got)
			}
		})
	}
}

func TestSumForward(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		a, b []int
		want []int
	}{
		{a: []int{}, b: []int{}, want: []int{0}},
		{a: []int{0}, b: []int{0}, want: []int{0}},
		{a: []int{1}, b: []int{2}, want: []int{3}},
		{a: []int{2, 7}, b: []int{3, 4, 1}, want: []int{3, 6, 8}},
		{a: []int{7, 1, 6}, b: []int{5, 9, 2}, want: []int{1, 3, 0, 8}},
	}

	for _, test := range subtests {
		test := test

		a := ch02.NewLinkedList(test.a)
		b := ch02.NewLinkedList(test.b)
		desc := fmt.Sprintf("%s %s", a, b)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			got := ch02.SumForward(a, b)
			want := ch02.NewLinkedList(test.want)

			if got.String() != want.String() {
				t.Errorf("\nwant: %s\n got: %s", want, got)
			}
		})
	}
}
