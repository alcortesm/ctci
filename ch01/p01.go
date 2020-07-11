package ch01

func IsUnique(s string) bool {
	seen := make(map[rune]struct{})

	for _, r := range s {
		if _, ok := seen[r]; ok {
			return false
		}

		seen[r] = struct{}{}
	}

	return true
}

func IsUniqueNoDataStructs(s string) bool {
	runes := []rune(s)

	for i, r := range runes {
		for _, o := range runes[i+1:] {
			if r == o {
				return false
			}
		}
	}

	return true
}
