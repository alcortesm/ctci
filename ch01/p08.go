package ch01

func ZeroInPlace(m [][]int32) {
	n := len(m)

	firstRowHasZero := hasFirstRowAZero(m)
	firstColHasZero := hasFirstColAZero(m)

	// find zeros on r,c > 0 and keeps track of them in the
	// first row and col
	for r := 1; r < n; r++ {
		for c := 1; c < n; c++ {
			if m[r][c] == 0 {
				m[0][c] = 0
				m[r][0] = 0
			}
		}
	}

	// null all rows > 0 with zeros in its first column
	for r := 1; r < n; r++ {
		if m[r][0] == 0 {
			for c := 1; c < n; c++ {
				m[r][c] = 0
			}
		}
	}

	// null all columns > 0 with zeros in its first column
	for c := 1; c < n; c++ {
		if m[0][c] == 0 {
			for r := 1; r < n; r++ {
				m[r][c] = 0
			}
		}
	}

	// null first row if needed
	if firstRowHasZero {
		for c := 0; c < n; c++ {
			m[0][c] = 0
		}
	}

	// null first column if needed
	if firstColHasZero {
		for r := 0; r < n; r++ {
			m[r][0] = 0
		}
	}

	return
}

func hasFirstRowAZero(m [][]int32) bool {
	n := len(m)

	for c := 0; c < n; c++ {
		if m[0][c] == 0 {
			return true
		}
	}

	return false
}

func hasFirstColAZero(m [][]int32) bool {
	n := len(m)

	for r := 0; r < n; r++ {
		if m[r][0] == 0 {
			return true
		}
	}

	return false
}
