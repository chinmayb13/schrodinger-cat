package graphs4

/*
Given a 2-D board A of size N x M containing 'X' and 'O', capture all regions surrounded by 'X'.
A region is captured by flipping all 'O's into 'X's in that surrounded region.

Problem Constraints
1 <= N, M <= 1000

Input Format
First and only argument is a N x M character matrix A.

Output Format
Make changes to the the input only as matrix is passed by reference.

Example Input
Input 1:
A = [

	  [X, X, X, X],
	  [X, O, O, X],
	  [X, X, O, X],
	  [X, O, X, X]
	]

Input 2:
A = [

	  [X, O, O],
	  [X, O, X],
	  [O, O, O]
	]

Example Output
Output 1:
After running your function, the board should be:
A = [

	  [X, X, X, X],
	  [X, X, X, X],
	  [X, X, X, X],
	  [X, O, X, X]
	]

Output 2:
After running your function, the board should be:
A = [

	  [X, O, O],
	  [X, O, X],
	  [O, O, O]
	]

Example Explanation
Explanation 1:
O in (4,2) is not surrounded by X from below.

Explanation 2:
No O's are surrounded.
*/
func GetCapturedRegion(board [][]byte) {
	N := len(board)
	M := len(board[0])
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	//check the boundaries at columns containing O
	// if yes, mark the connected component as M
	for i := 0; i < N; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0, N, M, board, dir)
		}

		if board[i][M-1] == 'O' {
			dfs(i, M-1, N, M, board, dir)
		}
	}

	//check the boundaries at rows containing O
	// if yes, mark the connected component as M
	for i := 0; i < M; i++ {
		if board[0][i] == 'O' {
			dfs(0, i, N, M, board, dir)
		}

		if board[N-1][i] == 'O' {
			dfs(N-1, i, N, M, board, dir)
		}
	}

	//convert M to O and O to X
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'M' {
				board[i][j] = 'O'
			}
		}
	}

}

func dfs(row, col, N, M int, board [][]byte, dir [][]int) {
	board[row][col] = 'M'
	for i := range dir {
		newRow := row + dir[i][0]
		newCol := col + dir[i][1]
		if check(newRow, newCol, N, M, board) {
			dfs(newRow, newCol, N, M, board, dir)
		}
	}
}

func check(row, col, N, M int, board [][]byte) bool {
	return row >= 0 && row < N && col >= 0 && col < M && board[row][col] == 'O'
}
