package ch01_test

import (
	"testing"

	"github.com/alcortesm/ctci-6th/ch01"
)

func TestIsPalindromePermutation(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input string
		want  bool
	}{
		{input: "", want: true},

		{input: "a", want: true},
		{input: ".a ", want: true}, // ignore non-lowercase latin letters

		{input: "aza", want: true},
		{input: " a _/  a ", want: true},

		{input: "aaa", want: true},

		{input: "aba", want: true},
		{input: "aab", want: true},
		{input: "baa", want: true},

		{input: "abba", want: true},
		{input: "aabb", want: true},
		{input: "bbaa", want: true},
		{input: "abab", want: true},
		{input: "baba", want: true},
		{input: "bab a", want: true},

		{input: "ab", want: false},
		{input: "a.b", want: false},
		{input: "aaab", want: false},
		{input: "aaa  b", want: false},
		{input: "abcbad", want: false},
		{input: "a  bcb  ad", want: false},
	}

	for _, test := range subtests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			got := ch01.IsPalindromePermutation(test.input)

			if got != test.want {
				t.Fatalf("want %t, got %t", test.want, got)
			}
		})
	}
}
