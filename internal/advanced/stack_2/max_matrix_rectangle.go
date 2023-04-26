package stack2

/*
Given a 2D binary matrix of integers A containing 0's and 1's of size N x M.
Find the largest rectangle containing only 1's and return its area.
Note: Rows are numbered from top to bottom and columns are numbered from left to right.

Input Format
The only argument given is the integer matrix A.

Output Format
Return the area of the largest rectangle containing only 1's.

Constraints
1 <= N, M <= 1000
0 <= A[i] <= 1

For Example
Input 1:

	A = [   [0, 0, 1]
	        [0, 1, 1]
	        [1, 1, 1]   ]

Output 1:

	4

Input 2:

	A = [   [0, 1, 0, 1]
	        [1, 0, 1, 0]    ]

Output 2:

	1
*/
func GetMaxRectangleArea(A [][]int) int {
	var ans int
	//perform column wise sum to calculate the height of rectangle
	doColumnWiseSum(A)
	for i := range A {
		nsl := getNSL(A[i])
		nsr := getNSR(A[i])
		for j := range A[i] {
			ans = max(ans, A[i][j]*(nsr[j]-nsl[j]-1))
		}
	}
	return ans
}

func doColumnWiseSum(inp [][]int) {
	rows := len(inp)
	cols := len(inp[0])
	for c := 0; c < cols; c++ {
		sum := 0
		for r := 0; r < rows; r++ {
			if inp[r][c] > 0 {
				inp[r][c] += sum
				sum = inp[r][c]
			} else {
				sum = 0
			}
		}
	}
}
