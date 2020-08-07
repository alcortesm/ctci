package ch02_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci/ch02"
)

func TestKthToLast(t *testing.T) {
	t.Parallel()

	type input struct {
		list []int
		k    int
	}

	subtests := []struct {
		input input
		want  int
	}{
		{input: input{list: []int{0}, k: 0}, want: 0},

		{input: input{list: []int{0, 1}, k: 0}, want: 1},
		{input: input{list: []int{0, 1}, k: 1}, want: 0},

		{input: input{list: []int{0, 1, 2}, k: 0}, want: 2},
		{input: input{list: []int{0, 1, 2}, k: 1}, want: 1},
		{input: input{list: []int{0, 1, 2}, k: 2}, want: 0},

		{input: input{list: []int{0, 1, 2, 3}, k: 0}, want: 3},
		{input: input{list: []int{0, 1, 2, 3}, k: 1}, want: 2},
		{input: input{list: []int{0, 1, 2, 3}, k: 2}, want: 1},
		{input: input{list: []int{0, 1, 2, 3}, k: 3}, want: 0},
	}

	for _, test := range subtests {
		test := test

		l := ch02.NewLinkedList(test.input.list)
		desc := fmt.Sprintf("%q %d", l, test.input.k)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			got, err := ch02.KthToLast(l, test.input.k)
			if err != nil {
				t.Fatal(err)
			}

			if test.want != got {
				t.Errorf("want %d got %d", test.want, got)
			}
		})
	}
}
