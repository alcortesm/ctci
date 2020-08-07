package ch02_test

import (
	"testing"

	"github.com/alcortesm/ctci/ch02"
)

func TestHasLoop(t *testing.T) {
	t.Parallel()

	n := make([]ch02.Node, 23)
	for i := range n {
		n[i].Data = i
	}

	l0 := ch02.LinkedList{First: &n[0]}
	n[0].Next = &n[1]
	n[1].Next = &n[1]

	l1 := ch02.LinkedList{First: &n[2]}
	n[2].Next = &n[3]
	n[3].Next = &n[4]
	n[4].Next = &n[3]

	l2 := ch02.LinkedList{First: &n[5]}
	n[5].Next = &n[6]
	n[6].Next = &n[7]
	n[7].Next = &n[8]
	n[8].Next = &n[6]

	l3 := ch02.LinkedList{First: &n[9]}
	n[9].Next = &n[10]
	n[10].Next = &n[11]
	n[11].Next = &n[12]
	n[12].Next = &n[13]
	n[13].Next = &n[14]
	n[14].Next = &n[15]
	n[15].Next = &n[16]
	n[16].Next = &n[17]
	n[17].Next = &n[18]
	n[18].Next = &n[19]
	n[19].Next = &n[12]

	l4 := ch02.LinkedList{First: &n[20]}
	n[20].Next = &n[21]
	n[21].Next = &n[22]

	subtests := []struct {
		desc  string
		input ch02.LinkedList
		want  *ch02.Node
	}{
		{desc: "0 1 1", input: l0, want: &n[1]},
		{desc: "2 3 4 3", input: l1, want: &n[3]},
		{desc: "5 6 7 8 6", input: l2, want: &n[6]},
		{desc: "9 10 11 12 13 14 15 16 17 18 19 11", input: l3, want: &n[12]},
		{desc: "20 21 22", input: l4, want: nil},
	}

	for _, test := range subtests {
		test := test

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			got := ch02.FirstNodeInLoop(test.input)
			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
