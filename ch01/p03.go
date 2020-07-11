package ch01

func URLify(runes []rune, n int) []rune {
	slow := len(runes) - 1

	for i := n - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			runes[slow-2] = '%'
			runes[slow-1] = '2'
			runes[slow] = '0'
			slow -= 3
		} else {
			runes[slow] = runes[i]
			slow--
		}
	}

	return runes
}
