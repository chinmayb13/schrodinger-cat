package dp2

/*
Problem Description
Given a grid of size n * m, lets assume you are starting at (1,1) and your goal is to reach (n, m). At any instance, if you are on (x, y), you can either go to (x, y + 1) or (x + 1, y).

Now consider if some obstacles are added to the grids. How many unique paths would there be? An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Problem Constraints
1 <= n, m <= 100
A[i][j] = 0 or 1

Input Format
Firts and only argument A is a 2D array of size n * m.

Output Format
Return an integer denoting the number of unique paths from (1, 1) to (n, m).

Example Input
Input 1:
A = [
       [0, 0, 0]
       [0, 1, 0]
       [0, 0, 0]
    ]

Input 2:
A = [
       [0, 0, 0]
       [1, 1, 1]
       [0, 0, 0]
    ]

Example Output
Output 1:
2

Output 2:
0

Example Explanation
Explanation 1:
Possible paths to reach (n, m): {(1, 1), (1, 2), (1, 3), (2, 3), (3, 3)} and {(1 ,1), (2, 1), (3, 1), (3, 2), (3, 3)}
So, the total number of unique paths is 2.

Explanation 2:
It is not possible to reach (n, m) from (1, 1). So, ans is 0.
*/
func UniquePathsWithObstacles(A [][]int) int {
	if A[0][0] == 1 || A[len(A)-1][len(A[0])-1] == 1 {
		return 0
	}
	rows, cols := len(A), len(A[0])
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return countWaysRecursive(rows-1, cols-1, A, memo)
}

func countWays(inp [][]int) int {
	rows, cols := len(inp), len(inp[0])
	//create two arrays
	cur := make([]int, cols)
	prev := make([]int, cols)
	for i := 0; i < cols; i++ {
		//if there is blocker i.e value is 1, stop
		if inp[0][i] == 1 {
			break
		}
		cur[i] = 1
	}
	//swap cur, prev
	prev, cur = cur, prev

	//iterate over rows
	for i := 1; i < rows; i++ {
		//if the 0th index of prev array is 1 and ith row is 0, save 1, else save 0
		if prev[0] == 1 && inp[i][0] == 0 {
			cur[0] = 1
		} else {
			cur[0] = 0
		}

		//traverse cols
		for j := 1; j < cols; j++ {
			//cur cell is sum of (cur row and prev col) and (prev row and cur col)
			cur[j] = cur[j-1] + prev[j]

			//if the inp cell is a blocker, save 0
			if inp[i][j] == 1 {
				cur[j] = 0
			}

		}
		//swap
		prev, cur = cur, prev
	}
	return prev[cols-1]
}

func countWaysRecursive(row, col int, inp [][]int, memo [][]int) int {
	//base condition
	if row < 0 || col < 0 {
		return 0
	}
	if inp[row][col] == 1 {
		return 0
	}
	if row == 0 && col == 0 {
		return 1
	}
	//we can reach a point either from left or from top
	if memo[row][col] == -1 {
		memo[row][col] = countWaysRecursive(row-1, col, inp, memo) + countWaysRecursive(row, col-1, inp, memo)
	}
	return memo[row][col]
}
