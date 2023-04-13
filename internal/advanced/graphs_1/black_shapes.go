package graphs1

/*
Problem Description

Given character matrix A of O's and X's, where O = white, X = black.
Return the number of black shapes. A black shape consists of one or more adjacent X's (diagonals not included)

Problem Constraints
1 <= |A|,|A[0]| <= 1000
A[i][j] = 'X' or 'O'

Input Format
The First and only argument is character matrix A.

Output Format
Return a single integer denoting number of black shapes.

Example Input
Input 1:
A = [ [X, X, X], [X, X, X], [X, X, X] ]

Input 2:
A = [ [X, O], [O, X] ]

Example Output
Output 1:
1

Output 2:
2

Example Explanation
Explanation 1:
All X's belong to single shapes

Explanation 2:
Both X's belong to different shapes
*/
func CountBlackShapes(A []string) int {
	inp := A
	rows := len(inp)
	cols := len(inp[0])

	//create visited array
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	//create direction array
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			//if a valid cell is unvisited, do DFS
			if inp[i][j] == 'X' && !visited[i][j] {
				count++
				depthSearch(i, j, inp, visited, dir)
			}
		}
	}
	return count
}

func depthSearch(row, col int, inp []string, visited [][]bool, dir [][]int) {
	rows := len(inp)
	cols := len(inp[0])
	visited[row][col] = true
	//check in four directions
	for k := 0; k < 4; k++ {
		newRow := row + dir[k][0]
		newCol := col + dir[k][1]
		//if cell is valid and unvisited, do dfs
		if isValidCell(newRow, newCol, rows, cols, inp, visited) {
			depthSearch(newRow, newCol, inp, visited, dir)
		}
	}
}

func isValidCell(row, col, rows, cols int, inp []string, visited [][]bool) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols && inp[row][col] == 'X' && !visited[row][col]
}
