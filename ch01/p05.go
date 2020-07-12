package ch01

func OneAway(a, b []rune) bool {
	switch {
	case len(a) == len(b):
		return oneAwayReplace(a, b)
	case len(a)+1 == len(b):
		return oneAwayInsert(a, b)
	case len(a)-1 == len(b): // one delete:
		return oneAwayInsert(b, a) // as insert but input args reversed
	default:
		return false
	}
}

// returns if the strings only differ in one rune
func oneAwayReplace(a, b []rune) bool {
	nDiffs := 0

	for i := range a {
		if a[i] != b[i] {
			nDiffs++
			if nDiffs > 1 {
				return false
			}
		}
	}

	return true
}

// returns if the strings only differ in one inserted rune
func oneAwayInsert(a, b []rune) bool {
	foundFirstDiff := false

	for i, j := 0, 0; i < len(a); i, j = i+1, j+1 {
		if a[i] != b[j] {
			if foundFirstDiff {
				return false
			}

			foundFirstDiff = true
			// cancel the autoincrement in i, this is, check the same
			// rune in a again
			i--
		}
	}

	return true
}
