package hashing2

/*
Determine if a Sudoku is valid, according to: http://sudoku.com.au/TheRules.aspx
The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

The input corresponding to the above configuration :

["53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"]
A partially filled sudoku which is valid.

Note:

A valid Sudoku board (partially filled) is not necessarily solvable. Only the filled cells need to be validated.
Return 0 / 1 ( 0 for false, 1 for true ) for this problem
*/

//Approach: Idea is to create map for each row (9 rows in total), each column(9 columns in total), each cube(9 cubes in total)
//and validate if the element present in any of these is being repeated
func IsValidSudoku(A []string) int {
	rowSet := make([]map[byte]struct{}, 9)
	for i := range rowSet {
		rowSet[i] = make(map[byte]struct{})
	}
	colSet := make([]map[byte]struct{}, 9)
	for i := range colSet {
		colSet[i] = make(map[byte]struct{})
	}
	cubeSet := make([][]map[byte]struct{}, 3)
	for i := range cubeSet {
		cubeSet[i] = make([]map[byte]struct{}, 3)
		for j := range cubeSet[i] {
			cubeSet[i][j] = make(map[byte]struct{})
		}
	}

	for i := range A {
		for j := range A[i] {
			if A[i][j] != '.' {
				_, rowDuplicate := rowSet[i][A[i][j]]
				_, columnDuplicate := colSet[j][A[i][j]]
				_, cubeDuplicate := cubeSet[i/3][j/3][A[i][j]]
				if rowDuplicate || columnDuplicate || cubeDuplicate {
					return 0
				}
				rowSet[i][A[i][j]] = struct{}{}
				colSet[j][A[i][j]] = struct{}{}
				cubeSet[i/3][j/3][A[i][j]] = struct{}{}
			}
		}
	}
	return 1

}
