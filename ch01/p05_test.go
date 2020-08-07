package ch01_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci/ch01"
)

func TestOneAway(t *testing.T) {
	t.Parallel()

	type input struct {
		a, b string
	}

	subtests := []struct {
		input input
		want  bool
	}{
		// empty
		{input: input{"", ""}, want: true},

		// same
		{input: input{"a", "a"}, want: true},
		{input: input{"abcd", "abcd"}, want: true},

		// one remove away
		{input: input{"a", ""}, want: true},
		{input: input{"ab", "a"}, want: true},
		{input: input{"ab", "b"}, want: true},

		// one insert away
		{input: input{"", "a"}, want: true},
		{input: input{"a", "ab"}, want: true},
		{input: input{"a", "ba"}, want: true},

		// one replace away
		{input: input{"a", "d"}, want: true},
		{input: input{"ab", "db"}, want: true},
		{input: input{"ab", "ad"}, want: true},

		// more than one edit away
		// two removes
		{input: input{"ab", ""}, want: false},
		{input: input{"abcd", "ac"}, want: false},
		// two inserts
		{input: input{"", "ab"}, want: false},
		{input: input{"ac", "abcd"}, want: false},
		// two replaces
		{input: input{"ab", "cd"}, want: false},
		{input: input{"abcd", "accb"}, want: false},
		// replace + insert
		{input: input{"a", "dd"}, want: false},
		// insert + insert
		{input: input{"ab", "dabd"}, want: false},
		// insert + remove (or two replaces)
		{input: input{"abc", "dac"}, want: false},
	}

	for _, test := range subtests {
		test := test

		desc := fmt.Sprintf("%s-%s", test.input.a, test.input.b)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			a, b := []rune(test.input.a), []rune(test.input.b)

			got := ch01.OneAway(a, b)

			if got != test.want {
				t.Fatalf("want %t, got %t", test.want, got)
			}

			// same test but in reversing the order of the strings
			got = ch01.OneAway(b, a)

			if got != test.want {
				t.Fatalf("(reverse) want %t, got %t", test.want, got)
			}
		})
	}
}
