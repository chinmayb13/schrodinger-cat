package graphs1

/*
Problem Description
There is a rectangle with left bottom as (0, 0) and right up as (x, y).

There are N circles such that their centers are inside the rectangle.

Radius of each circle is R. Now we need to find out if it is possible that we can move from (0, 0) to (x, y) without touching any circle.

Note : We can move from any cell to any of its 8 adjecent neighbours and we cannot move outside the boundary of the rectangle at any point of time.

Problem Constraints
0 <= x , y, R <= 100
1 <= N <= 1000

Center of each circle would lie within the grid

Input Format
1st argument given is an Integer x , denoted by A in input.
2nd argument given is an Integer y, denoted by B in input.
3rd argument given is an Integer N, number of circles, denoted by C in input.
4th argument given is an Integer R, radius of each circle, denoted by D in input.
5th argument given is an Array A of size N, denoted by E in input, where A[i] = x cordinate of ith circle
6th argument given is an Array B of size N, denoted by F in input, where B[i] = y cordinate of ith circle

Output Format
Return YES or NO depending on weather it is possible to reach cell (x,y) or not starting from (0,0).

Example Input
Input 1:
x = 2
y = 3
N = 1
R = 1
A = [2]
B = [3]

Input 2:
x = 1
y = 1
N = 1
R = 1
A = [1]
B = [1]

Example Output
Output 1:
NO

Output 2:
NO

Example Explanation
Explanation 1:
There is NO valid path in this case

Explanation 2:
There is NO valid path in this case
*/
func DoesPathExist(A int, B int, C int, D int, E []int, F []int) string {
	rows := A + 1
	cols := B + 1
	radius := D
	N := C
	matrix := make([][]bool, rows)
	visited := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]bool, cols)
		visited[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for k := 0; k < N; k++ {
				//set circle touching coordinates to true
				if getSqrDist(i-E[k], j-F[k]) <= radius*radius {
					matrix[i][j] = true
				}
			}
		}
	}

	//if source and destination are touched by circle, return NO
	if matrix[0][0] || matrix[rows-1][cols-1] {
		return "NO"
	}
	dir := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	//check the path using DFS
	if pathExists(0, 0, rows, cols, matrix, visited, dir) {
		return "YES"
	}
	return "NO"
}

func getSqrDist(inp1, inp2 int) int {
	return inp1*inp1 + inp2*inp2
}

func validateCell(row, col, rows, cols int, matrix, visited [][]bool) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols && /*matrix cell should be false to be valid*/ !matrix[row][col] && !visited[row][col]
}

func pathExists(row, col, rows, cols int, matrix, visited [][]bool, dir [][]int) bool {
	if row == rows-1 && col == cols-1 {
		return true
	}
	visited[row][col] = true
	for i := 0; i < 8; i++ {
		newRow := row + dir[i][0]
		newCol := col + dir[i][1]
		//validate adjacent cell
		if validateCell(newRow, newCol, rows, cols, matrix, visited) {
			if pathExists(newRow, newCol, rows, cols, matrix, visited, dir) {
				return true
			}
		}
	}
	return false
}
