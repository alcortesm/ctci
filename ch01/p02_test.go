package ch01_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci-6th/ch01"
)

func TestIsPermutation(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input [2]string
		want  bool
	}{
		{input: [2]string{"", ""}, want: true},
		{input: [2]string{"a", "a"}, want: true},
		{input: [2]string{"a", ""}, want: false},
		{input: [2]string{"ab", "ab"}, want: true},
		{input: [2]string{"ab", "ba"}, want: true},
		{input: [2]string{"ab", "cc"}, want: false},
		{input: [2]string{"ab", "ac"}, want: false},
	}

	for _, test := range subtests {
		test := test

		desc := fmt.Sprintf("%q", test.input)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			got := ch01.ArePermutation(test.input[0], test.input[1])

			if got != test.want {
				t.Fatalf("want %t, got %t", test.want, got)
			}

			// ArePermutation is commutative
			got = ch01.ArePermutation(test.input[1], test.input[0])

			if got != test.want {
				t.Fatalf("(commutative) want %t, got %t", test.want, got)
			}
		})
	}
}
