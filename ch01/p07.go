package ch01

func RotateInPlace(m [][]int32) {
	n := len(m)

	// we are going to rotate the matrix layer by layer.
	// A layer is a ring of matrix elements with a constant distance
	// from its center:
	//
	// 0 1 2 3
	// 4 5 6 7
	// 8 9 0 a
	// b c d e
	//
	// layer 0 is 01237aedcb84
	// layer 1 is 5609
	//
	// there are n/2 layer in a matrix
	for layer := 0; layer < n/2; layer++ {
		// length of the sides of the layer
		side := n - (2 * layer)
		// index of the top row of the layer
		topRow := layer
		// index of the bottom row of the layer
		bottomRow := layer + side - 1
		// index of the left column of the layer
		leftColumn := layer
		// index of the right column of the layer
		rightColumn := layer + side - 1

		// each layer side has "side" elements, we are going to move
		// them one by one, from left to top, bottom to left,
		// right to bottom and top to right.
		for i := 0; i < side-1; i++ {
			// save top element at column i for later
			top := m[topRow][leftColumn+i]
			// copy ith element from left column to top row
			m[topRow][leftColumn+i] = m[bottomRow-i][leftColumn]
			// copy ith element from bottom row to left column
			m[bottomRow-i][leftColumn] = m[bottomRow][rightColumn-i]
			// copy ith element from right column to bottom row
			m[bottomRow][rightColumn-i] = m[topRow+i][rightColumn]
			// copy ith element from top row to right column
			m[topRow+i][rightColumn] = top
		}
	}
}
