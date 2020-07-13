package ch01_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci-6th/ch01"
)

func TestIsStringRotation(t *testing.T) {
	t.Parallel()

	type input struct {
		a, b string
	}

	subtests := []struct {
		input input
		want  bool
	}{
		{input: input{"", ""}, want: true},

		{input: input{"a", "a"}, want: true},
		{input: input{"a", "b"}, want: false},

		{input: input{"", "a"}, want: false},
		{input: input{"a", ""}, want: false},

		{input: input{"ab", "ba"}, want: true},
		{input: input{"ab", "ac"}, want: false},

		{input: input{"abc", "abc"}, want: true},
		{input: input{"abc", "cab"}, want: true},
		{input: input{"abc", "bca"}, want: true},
		{input: input{"abc", "acb"}, want: false},

		{input: input{"erbottlewat", "waterbottle"}, want: true},
	}

	for _, test := range subtests {
		test := test

		desc := fmt.Sprintf("%s %s", test.input.a, test.input.b)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			got := ch01.IsStringRotation(test.input.a, test.input.b)

			if got != test.want {
				t.Fatalf("want %t, got %t", test.want, got)
			}
		})
	}
}
