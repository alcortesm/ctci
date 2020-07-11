package ch01_test

import (
	"testing"

	"github.com/alcortesm/ctci-6th/ch01"
)

func TestIsUnique(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input string
		want  bool
	}{
		{input: "", want: true},
		{input: "a", want: true},
		{input: "aa", want: false},
		{input: "ab", want: true},
		{input: "áí", want: true},
		{input: "áá", want: false},
		{input: "alberto", want: true},
		{input: "alberto e", want: false},
		{input: "alberto é", want: true},
	}

	for _, test := range subtests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			got := ch01.IsUnique(test.input)

			if got != test.want {
				t.Fatalf("want %t, got %t", test.want, got)
			}
		})
	}
}

func TestIsUniqueNoDataStructs(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		input string
		want  bool
	}{
		{input: "", want: true},
		{input: "a", want: true},
		{input: "aa", want: false},
		{input: "ab", want: true},
		{input: "áá", want: false},
		{input: "alberto", want: true},
		{input: "alberto e", want: false},
		{input: "alberto é", want: true},
	}

	for _, test := range subtests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			got := ch01.IsUniqueNoDataStructs(test.input)

			if got != test.want {
				t.Fatalf("want %t, got %t", test.want, got)
			}
		})
	}
}
