package ch01_test

import (
	"testing"

	"github.com/alcortesm/ctci/ch01"
	"github.com/google/go-cmp/cmp"
)

func TestRotate(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		desc  string
		input [][]int32
		want  [][]int32
	}{
		{
			desc: "0",
			input: [][]int32{
				{},
			},
			want: [][]int32{
				{},
			},
		},
		{
			desc: "1",
			input: [][]int32{
				{1},
			},
			want: [][]int32{
				{1},
			},
		},
		{
			desc: "2",
			input: [][]int32{
				{1, 2},
				{3, 4},
			},
			want: [][]int32{
				{3, 1},
				{4, 2},
			},
		},
		{
			desc: "3",
			input: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int32{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			desc: "4",
			input: [][]int32{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]int32{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
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

			ch01.RotateInPlace(dup)

			if diff := cmp.Diff(test.want, dup); diff != "" {
				t.Fatalf("(-want +got):\n%s", diff)
			}
		})
	}
}
