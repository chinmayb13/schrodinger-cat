package arrays2

import "math"

/*
Problem Description
Given a 2D integer matrix A of size N x N, find a B x B submatrix where B<= N and B>= 1, such that the sum of all the elements in the submatrix is maximum.

Problem Constraints
1 <= N <= 103.
1 <= B <= N
-102 <= A[i][j] <= 102.

Input Format
First arguement is an 2D integer matrix A.
Second argument is an integer B.

Output Format
Return a single integer denoting the maximum sum of submatrix of size B x B.

Example Input
Input 1:

	A = [
	       [1, 1, 1, 1, 1]
	       [2, 2, 2, 2, 2]
	       [3, 8, 6, 7, 3]
	       [4, 4, 4, 4, 4]
	       [5, 5, 5, 5, 5]
	    ]

B = 3

Input 2:

	A = [
	       [2, 2]
	       [2, 2]
	   ]
	B = 2

Example Output
Output 1:
48

Output 2:
8

Example Explanation
Explanation 1:
Maximum sum 3 x 3 matrix is
8 6 7
4 4 4
5 5 5
Sum = 48

Explanation 2:
Maximum sum 2 x 2 matrix is
2 2
2 2
Sum = 8
*/
func GetMaxSquareSubMatrixSum(A [][]int, B int) int {
	maxSum := math.MinInt
	buildPrefixMatrixSum(A)
	for r := B - 1; r < len(A); r++ {
		for c := B - 1; c < len(A[0]); c++ {
			sum := A[r][c]
			if (r + c) > (2*B - 2) {
				if r == (B - 1) {
					sum -= A[r][c-B]
				} else if c == (B - 1) {
					sum -= A[r-B][c]
				} else {
					sum = sum - (A[r][c-B] + A[r-B][c] - A[r-B][c-B])
				}
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func buildPrefixMatrixSum(A [][]int) {
	for i := range A {
		for j := range A[i] {
			if i+j > 0 {
				if i == 0 {
					A[i][j] += A[i][j-1]
				} else if j == 0 {
					A[i][j] += A[i-1][j]
				} else {
					A[i][j] += (A[i][j-1] + A[i-1][j] - A[i-1][j-1])
				}
			}
		}
	}
}
