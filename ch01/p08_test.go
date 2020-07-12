package ch01_test

import (
	"testing"

	"github.com/alcortesm/ctci-6th/ch01"
	"github.com/google/go-cmp/cmp"
)

func TestZeroInPlace(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		desc  string
		input [][]int32
		want  [][]int32
	}{
		{
			desc:  "0",
			input: [][]int32{},
			want:  [][]int32{},
		},
		{
			desc: "1 no zeros",
			input: [][]int32{
				{1},
			},
			want: [][]int32{
				{1},
			},
		},
		{
			desc: "1 with zero",
			input: [][]int32{
				{0},
			},
			want: [][]int32{
				{0},
			},
		},
		{
			desc: "2 with no zeros",
			input: [][]int32{
				{1, 2},
				{3, 4},
			},
			want: [][]int32{
				{1, 2},
				{3, 4},
			},
		},
		{
			desc: "2 with 1 zero",
			input: [][]int32{
				{1, 0},
				{3, 4},
			},
			want: [][]int32{
				{0, 0},
				{3, 0},
			},
		},
		{
			desc: "2 with 2 zeros",
			input: [][]int32{
				{1, 2},
				{0, 0},
			},
			want: [][]int32{
				{0, 0},
				{0, 0},
			},
		},
		{
			desc: "3 no zeros",
			input: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			desc: "3 one side zero",
			input: [][]int32{
				{1, 2, 3},
				{4, 5, 0},
				{7, 8, 9},
			},
			want: [][]int32{
				{1, 2, 0},
				{0, 0, 0},
				{7, 8, 0},
			},
		},
		{
			desc: "3 one corner zero",
			input: [][]int32{
				{1, 2, 0},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int32{
				{0, 0, 0},
				{4, 5, 0},
				{7, 8, 0},
			},
		},
		{
			desc: "3 one central zero",
			input: [][]int32{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			want: [][]int32{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		{
			desc: "3 one two zeros",
			input: [][]int32{
				{1, 2, 0},
				{4, 0, 6},
				{7, 8, 9},
			},
			want: [][]int32{
				{0, 0, 0},
				{0, 0, 0},
				{7, 0, 0},
			},
		},
		{
			desc: "3 all zeros",
			input: [][]int32{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: [][]int32{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}

	for _, test := range subtests {
		test := test

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			dup := make([][]int32, len(test.input))
			for i, row := range test.input {
				dupRow := make([]int32, len(row))
				copy(dupRow, row)
				dup[i] = dupRow
			}

			ch01.ZeroInPlace(dup)

			if diff := cmp.Diff(test.want, dup); diff != "" {
				t.Fatalf("(-want +got):\n%s", diff)
			}
		})
	}
}
