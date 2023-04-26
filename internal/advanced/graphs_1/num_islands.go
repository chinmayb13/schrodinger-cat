package graphs1

/*
Problem Description
Given a matrix of integers A of size N x M consisting of 0 and 1. A group of connected 1's forms an island. From a cell (i, j) such that A[i][j] = 1 you can visit any cell that shares a corner with (i, j) and value in that cell is 1.

More formally, from any cell (i, j) if A[i][j] = 1 you can visit:

(i-1, j) if (i-1, j) is inside the matrix and A[i-1][j] = 1.
(i, j-1) if (i, j-1) is inside the matrix and A[i][j-1] = 1.
(i+1, j) if (i+1, j) is inside the matrix and A[i+1][j] = 1.
(i, j+1) if (i, j+1) is inside the matrix and A[i][j+1] = 1.
(i-1, j-1) if (i-1, j-1) is inside the matrix and A[i-1][j-1] = 1.
(i+1, j+1) if (i+1, j+1) is inside the matrix and A[i+1][j+1] = 1.
(i-1, j+1) if (i-1, j+1) is inside the matrix and A[i-1][j+1] = 1.
(i+1, j-1) if (i+1, j-1) is inside the matrix and A[i+1][j-1] = 1.
Return the number of islands.

NOTE: Rows are numbered from top to bottom and columns are numbered from left to right.

Problem Constraints
1 <= N, M <= 100
0 <= A[i] <= 1

Input Format
The only argument given is the integer matrix A.

Output Format
Return the number of islands.

Example Input
Input 1:
A = [

	  [0, 1, 0]
	  [0, 0, 1]
	  [1, 0, 0]
	]

Input 2:
A = [

	  [1, 1, 0, 0, 0]
	  [0, 1, 0, 0, 0]
	  [1, 0, 0, 1, 1]
	  [0, 0, 0, 0, 0]
	  [1, 0, 1, 0, 1]
	]

Example Output
Output 1:
2

Output 2:
5

Example Explanation
Explanation 1:
The 1's at position A[0][1] and A[1][2] forms one island.
Other is formed by A[2][0].

Explanation 2:
There 5 island in total.
*/
func GetNumIslands(A [][]int) int {
	inp := A
	N := len(inp)
	M := len(inp[0])
	//visited array to see if element has alreasy been visited or not
	visited := make([][]bool, N)
	for i := range visited {
		visited[i] = make([]bool, M)
	}
	//no of connected islands
	count := 0

	//direction array to move in all directions
	direction := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}, {1, -1}, {-1, 1}, {-1, -1}, {1, 1}}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			//if valid element, do DFS recursively
			if check(i, j, N, M, inp, visited) {
				//increase the counter
				count++
				doDFS(i, j, N, M, inp, direction, visited)
			}
		}
	}
	return count
}

func check(row, col, N, M int, inp [][]int, visited [][]bool) bool {
	return row >= 0 && row < N && col >= 0 && col < M && inp[row][col] == 1 && !visited[row][col]
}

func doDFS(row, col int, N, M int, inp [][]int, direction [][]int, visited [][]bool) {
	visited[row][col] = true
	//check in all directions, move recursively
	for k := 0; k < 8; k++ {
		newRow := row + direction[k][0]
		newCol := col + direction[k][1]
		if check(newRow, newCol, N, M, inp, visited) {
			doDFS(newRow, newCol, N, M, inp, direction, visited)
		}
	}
}
