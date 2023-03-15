package backtracking2

import "sort"

/*
Problem Description
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer A, return all distinct solutions to the n-queens puzzle.
Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
The final list should be generated in such a way that the indices of the queens in each list should be in reverse lexicographical order.

Problem Constraints
1 <= A <= 10

Input Format
First argument is an integer n denoting the size of chessboard

Output Format
Return an array consisting of all distinct solutions in which each element is a 2d char array representing a unique solution.

Example Input
Input 1:
A = 4

Input 2:
A = 1

Example Output
Output 1:
[
[".Q..",  // Solution 1
 "...Q",
 "Q...",
 "..Q."],

["..Q.",  // Solution 2
 "Q...",
 "...Q",
 ".Q.."]
]

Output 1:
[
 [Q]
]

Example Explanation
Explanation 1:
There exist only two distinct solutions to the 4-queens puzzle:

Explanation 1:
There exist only one distinct solutions to the 1-queens puzzle:
*/
func SolveNQueens(n int) [][]string {
	chess := make([][]byte, n)

	//intialize the chess board with all .
	for i := range chess {
		chess[i] = make([]byte, n)
		for j := range chess[i] {
			chess[i][j] = '.'
		}
	}

	//colMap to look up
	colMap := make(map[int]bool)

	//diagonalMap to look rightup
	//all diagonal element have same coordinate sum,
	//i+j=constant
	diagonalMap := make(map[int]bool)

	//antiDiagonal map to look left up
	//all anti diagonal element have same coordinate difference,
	//i-j=constant
	antiDiagonalMap := make(map[int]bool)
	var ans [][]string
	placeQueens(0, n, chess, colMap, diagonalMap, antiDiagonalMap, &ans)
	sort.Slice(ans, func(i, j int) bool {
		k := 0
		for k < 9 && ans[i][k] == ans[j][k] {
			k++
		}
		return ans[i][k] < ans[j][k]
	})
	return ans

}

func placeQueens(row, N int, chess [][]byte, colMap, diagonalMap, antiDiagonalMap map[int]bool, ans *[][]string) {
	if row == N {
		*ans = append(*ans, copyArr(chess))
		return
	}

	//loop on columns for a given row
	for i := range chess[row] {
		//if the current position's top, left up and right up, don't have any queens
		//place the queen
		if !colMap[i] && !diagonalMap[row+i] && !antiDiagonalMap[row-i] {
			chess[row][i] = 'Q'
			colMap[i] = true
			diagonalMap[row+i] = true
			antiDiagonalMap[row-i] = true
			placeQueens(row+1, N, chess, colMap, diagonalMap, antiDiagonalMap, ans)
			delete(antiDiagonalMap, row-i)
			delete(diagonalMap, row+i)
			delete(colMap, i)
			chess[row][i] = '.'
		}
	}
}

func copyArr(inp [][]byte) []string {
	var ans []string
	for i := range inp {
		ans = append(ans, string(inp[i]))
	}
	return ans
}
