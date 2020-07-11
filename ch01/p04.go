package ch01

func IsPalindromePermutation(s string) bool {
	// All runes in a Palindrome have an even frequency count, except
	// maybe for one of them (the central rune if the len is odd
	// contributes as a single rune).
	//
	// A sentence would be a permutation of a palindrome if it satisfy
	// that same condition, as you can always rearange its runes to
	// form a palindrome.
	v := bitVectorLetterAppearsOddNumberOfTimes(s)
	return isPowerOf2(v)
}

// returns a bit vector of whether each letter in the english 26
// letter alphabet appears and odd number of times in an string.
func bitVectorLetterAppearsOddNumberOfTimes(s string) uint32 {
	var result uint32 = 0

	for _, r := range s {
		// ignore non latin lowercase alphabet runes
		if r < 'a' || r > 'z' {
			continue
		}

		i := r - 'a'       // number of letter r in alphabet
		result ^= (1 << i) // toggle ith bit of v
	}

	return result
}

// returns true if n is 0 or a power of 2
func isPowerOf2(n uint32) bool {
	return n&(n-1) == 0
}
