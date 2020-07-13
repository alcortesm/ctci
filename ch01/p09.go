package ch01

import "strings"

func IsStringRotation(a, b string) bool {
	// strings.Contains(s, substr) reports whether substr is within s.
	// The exercise restricts the times we can use this function to only
	// one call.

	// early return and support for b being the empty string, which is
	// always a substring of any other string.
	if len(a) != len(b) {
		return false
	}

	double := a + a // last rune is not really needed, but it takes more effort to remove it than just copying it.

	return strings.Contains(double, b)
}
