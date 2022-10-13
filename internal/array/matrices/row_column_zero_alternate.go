package matrices

func DoRowColumnZeroAlternate(A [][]int) [][]int {
	rowZeroElement := false
	colZeroElement := false
	for r := range A {
		for c := range A[r] {
			if A[r][c] == 0 {
				//flag if the matrices contains a 0 in either 0th row or 0th column
				if r == 0 {
					rowZeroElement = true
				}
				if c == 0 {
					colZeroElement = true
				}
				//mark the row start and column start to 0
				A[0][c] = 0
				A[r][0] = 0
			}
		}
	}

	//skip the 0th row and 0th column for now
	for r := 1; r < len(A); r++ {
		for c := 1; c < len(A[0]); c++ {
			//check the row start and column start
			if A[0][c] == 0 || A[r][0] == 0 {

				A[r][c] = 0

			}
		}
	}

	// make the entire 0th row zero if flag is true
	if rowZeroElement {
		for c := 0; c < len(A[0]); c++ {
			A[0][c] = 0
		}
	}
	// make the entire 0th col zero if flag is true
	if colZeroElement {
		for r := 0; r < len(A); r++ {
			A[r][0] = 0
		}
	}

	return A
}
