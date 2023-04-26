package dp2

/*
Problem Description
Given a M x N grid A of integers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Return the minimum sum of the path.

NOTE: You can only move either down or right at any point in time.

Problem Constraints
1 <= M, N <= 2000
-1000 <= A[i][j] <= 1000

Input Format
First and only argument is a 2-D grid A.

Output Format
Return an integer denoting the minimum sum of the path.

Example Input
Input 1:
A = [

	  [1, 3, 2]
	  [4, 3, 1]
	  [5, 6, 1]
	]

Input 2:
A = [

	  [1, -3, 2]
	  [2, 5, 10]
	  [5, -5, 1]
	]

Example Output
Output 1:
8

Output 2:
-1

Example Explanation
Explanation 1:
The path will be: 1 -> 3 -> 2 -> 1 -> 1.

Input 2:
The path will be: 1 -> -3 -> 5 -> -5 -> 1.
*/
func MinPathSum(A [][]int) int {
	inp := A
	rows := len(inp)
	cols := len(inp[0])
	prev := make([]int, cols)
	cur := make([]int, cols)

	//prepopulate 0th row values
	prev[0] = inp[0][0]
	for col := 1; col < cols; col++ {
		prev[col] = prev[col-1] + inp[0][col]
	}

	//calculate cell values from row 1 onwards
	for row := 1; row < rows; row++ {
		//prepopulate 0th col
		cur[0] = prev[0] + inp[row][0]
		//calculate cell values from col 1 onwards
		for col := 1; col < cols; col++ {
			cur[col] = min(cur[col-1], prev[col]) + inp[row][col]
		}
		//swap arrays
		cur, prev = prev, cur
	}
	return prev[cols-1]

}
