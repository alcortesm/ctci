package ch01

import (
	"fmt"
)

func Compress(s string) string {
	runes := []rune(s)

	// very short strings are not compressed
	if len(runes) < 3 {
		return s
	}

	// only compress strings with valid runes
	for _, r := range runes {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}

		return s // don't compress if there are non-letters
	}

	// records how many times in a row the last rune has been seen
	// initialize it directly with the first rune in the input.
	lastBurst := struct {
		rune  rune
		times int
	}{

		rune:  runes[0],
		times: 1,
	}

	compressed := make([]rune, 0)

	// utility function to write the last burst in a compressed form
	compressLastBurst := func() {
		compressed = append(compressed, lastBurst.rune)
		if lastBurst.times > 1 {
			times := []rune(fmt.Sprintf("%d", lastBurst.times))
			compressed = append(compressed, times...)
		}
	}

	for _, r := range runes[1:] {
		// if the last burst continues, update times and continue
		if r == lastBurst.rune {
			lastBurst.times++
			continue
		}

		// otherwise the last burst has just ended and a new one begins
		compressLastBurst()
		lastBurst.rune = r
		lastBurst.times = 1
	}

	// compress the rune/burst at the end of the input
	compressLastBurst()

	// return the original if compressing didn't save any space
	if len(runes) == len(compressed) {
		return s
	}

	return string(compressed)
}
