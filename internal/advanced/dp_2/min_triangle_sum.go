package dp2

/*
Problem Description
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
Adjacent numbers for jth column of ith row is jth and (j+1)th column of (i + 1)th row

Problem Constraints
|A| <= 1000
A[i] <= 1000

Input Format
First and only argument is the vector of vector A defining the given triangle

Output Format
Return the minimum sum

Example Input
Input 1:
A = [

	     [2],
	    [3, 4],
	   [6, 5, 7],
	  [4, 1, 8, 3]
	]

Input 2:
A = [ [1] ]

Example Output
Output 1:
11

Output 2:
1

Example Explanation
Explanation 1:
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Explanation 2:
Only 2 can be collected.
*/
func MinimumTotal(triangle [][]int) int {
	memo := make([][]int, len(triangle))
	for i := range memo {
		memo[i] = make([]int, len(triangle[i]))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return getMinPathSum(0, 0, len(triangle), triangle, memo)
}

func getMinPathSum(row, col, rows int, inp [][]int, memo [][]int) int {
	if row == rows {
		return 0
	}
	if memo[row][col] == -1 {
		memo[row][col] = inp[row][col] + min(getMinPathSum(row+1, col, rows, inp, memo), getMinPathSum(row+1, col+1, rows, inp, memo))
	}
	return memo[row][col]
}

func getMinPathSumIterative(inp [][]int) int {
	rows := len(inp)
	memo := make([][]int, len(inp))
	for i := range memo {
		memo[i] = make([]int, len(inp[i]))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	//prepopulate last row elements
	for i := 0; i < len(memo[rows-1]); i++ {
		memo[rows-1][i] = inp[rows-1][i]
	}

	//iterate from second last row until 0
	for row := rows - 2; row >= 0; row-- {
		for col := 0; col < len(memo[row]); col++ {
			memo[row][col] = inp[row][col] + min(memo[row+1][col], memo[row+1][col+1])
		}
	}
	return memo[0][0]
}
