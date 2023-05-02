package graphs5

/*
Problem Description
Given a matrix of integers A of size N x N, where A[i][j] represents the weight of directed edge from i to j (i ---> j).
If i == j, A[i][j] = 0, and if there is no directed edge from vertex i to vertex j, A[i][j] = -1.
Return a matrix B of size N x N where B[i][j] = shortest path from vertex i to vertex j.
If there is no possible path from vertex i to vertex j , B[i][j] = -1

Note: Rows are numbered from top to bottom and columns are numbered from left to right.

Problem Constraints
1 <= N <= 200
-1 <= A[i][j] <= 1000000

Input Format
The first and only argument given is the integer matrix A.

Output Format
Return a matrix B of size N x N where B[i][j] = shortest path from vertex i to vertex j
If there is no possible path from vertex i to vertex j, B[i][j] = -1.

Example Input
A = [ [0 , 50 , 39]

	[-1 , 0 , 1]
	[-1 , 10 , 0] ]

Example Output
[ [0 , 49 , 39 ]

	[-1 , 0 , -1 ]
	[-1 , 10 , 0 ] ]

Example Explanation
Shortest Path from 1 to 2 would be 1 ---> 3 ---> 2 and not directly from 1 to 2,
All remaining distances remains same.
*/
func GetShortestPath(A [][]int) [][]int {
	N := len(A)
	for i := range A {
		for j := range A[i] {
			if A[i][j] == -1 {
				A[i][j] = 1e9
			}
		}
	}

	//check minimum distance using dynamic programming
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				//for each cell A[j][k] check if there exists a shorter path of
				//A[j][i]+A[i][k]
				if i != j && j != k && k != i {
					A[j][k] = min(A[j][k], A[j][i]+A[i][k])
				}
			}
		}
	}

	for i := range A {
		for j := range A[i] {
			if A[i][j] == 1e9 {
				A[i][j] = -1
			}
		}
	}
	return A

}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}
