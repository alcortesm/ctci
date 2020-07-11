package ch01

func ArePermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	// two strings are permutations of each other if they have the same
	// frequency count of runes

	fc := make(map[rune]int)

	for _, r := range a {
		fc[r]++
	}

	for _, r := range b {
		fc[r]--
	}

	for _, v := range fc {
		if v != 0 {
			return false
		}
	}

	return true
}
