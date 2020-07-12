package ch01_test

import (
	"testing"

	"github.com/alcortesm/ctci-6th/ch01"
)

func TestCompress(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input string
		want  string
	}{
		{input: "", want: ""},

		{input: "a", want: "a"},
		{input: "aa", want: "aa"},
		{input: "aaa", want: "a3"},

		{input: "aaab", want: "a3b"},
		{input: "aaabb", want: "a3b2"},
		{input: "aabbb", want: "a2b3"},

		{input: "abbbcdde", want: "ab3cd2e"},
	}

	for _, test := range subtests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			got := ch01.Compress(test.input)

			if got != test.want {
				t.Fatalf("want %q, got %q", test.want, got)
			}
		})
	}
}
