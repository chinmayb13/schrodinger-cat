package backtracking2

/*
Problem Description
Write a program to solve a Sudoku puzzle by filling the empty cells. Empty cells are indicated by the character '.' You may assume that there will be only one unique solution.

A sudoku puzzle,

and its solution numbers marked in red.

Problem Constraints
N = 9

Input Format
First argument is an array of array of characters representing the Sudoku puzzle.

Output Format
Modify the given input to the required answer.

Example Input
Input 1:

A = [[53..7....], [6..195...], [.98....6.], [8...6...3], [4..8.3..1], [7...2...6], [.6....28.], [...419..5], [....8..79]]


Example Output
Output 1:

[[534678912], [672195348], [198342567], [859761423], [426853791], [713924856], [961537284], [287419635], [345286179]]


Example Explanation
Explanation 1:

Look at the diagrams given in the question.
*/

//idea: maintain 9 integers each for rows, cols and grid
//for a position, see if k (1<=k<=9) can be inserted by checking
// if k+1 th bit is set in row, col and grid entry corresponding to the position

func SolveSudoku(board *[][]byte) {
	rows := make([]int, 9)
	cols := make([]int, 9)
	grid := make([][]int, 3)
	for i := range grid {
		grid[i] = make([]int, 3)
	}

	//for the provided input, fill the bits in the rows, cols and grid
	for i := range *board {
		for j := range (*board)[i] {
			val := (*board)[i][j]
			if val != '.' {
				num := 1 << int(val-'0')
				rows[i] |= num
				cols[j] |= num
				grid[i/3][j/3] |= num
			}
		}
	}
	placeNumber(0, 0, rows, cols, grid, board)
}

func placeNumber(row, col int, rows, cols []int, grid [][]int, board *[][]byte) bool {
	if row == 9 {
		return true
	}

	if (*board)[row][col] == '.' {
		for k := 1; k <= 9; k++ {
			val := 1 << k
			if (val&rows[row])+(val&cols[col])+(val&grid[row/3][col/3]) == 0 {
				(*board)[row][col] = byte(k + '0')
				rows[row] |= val
				cols[col] |= val
				grid[row/3][col/3] |= val
				if col == 8 {
					if placeNumber(row+1, 0, rows, cols, grid, board) {
						return true
					}
				} else {
					if placeNumber(row, col+1, rows, cols, grid, board) {
						return true
					}
				}
				grid[row/3][col/3] ^= val
				cols[col] ^= val
				rows[row] ^= val
				(*board)[row][col] = '.'
			}
		}
	} else {
		if col == 8 {
			return placeNumber(row+1, 0, rows, cols, grid, board)
		} else {
			return placeNumber(row, col+1, rows, cols, grid, board)
		}
	}
	return false
}
