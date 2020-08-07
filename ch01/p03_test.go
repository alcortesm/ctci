package ch01_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/ctci/ch01"
)

func TestURLify(t *testing.T) {
	t.Parallel()

	type input struct {
		runes string
		len   int
	}

	subtests := []struct {
		input input
		want  string
	}{
		{input: input{``, 0}, want: ``},
		{input: input{`a`, 1}, want: `a`},
		{input: input{`   `, 1}, want: `%20`},
		{input: input{` a  `, 2}, want: `%20a`},
		{input: input{`a   `, 2}, want: `a%20`},
		{input: input{`a b c    `, 5}, want: `a%20b%20c`},
	}

	for _, test := range subtests {
		test := test

		desc := fmt.Sprintf("%s %d", test.input.runes, test.input.len)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			got := ch01.URLify([]rune(test.input.runes), test.input.len)

			if string(got) != test.want {
				t.Fatalf("want %q, got %q", test.want, string(got))
			}
		})
	}
}
