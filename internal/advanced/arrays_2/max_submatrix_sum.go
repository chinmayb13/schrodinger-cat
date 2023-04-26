package arrays2

import "math"

/*
Problem Description
Given a row-wise and column-wise sorted matrix A of size N * M.
Return the maximum non-empty submatrix sum of this matrix.

Problem Constraints
1 <= N, M <= 1000
-109 <= A[i][j] <= 109

Input Format
The first argument is a 2D integer array A.

Output Format
Return a single integer that is the maximum non-empty submatrix sum of this matrix.

Example Input
Input 1:-

	-5 -4 -3

A = -1  2  3

	2  2  4

Input 2:-

	1 2 3

A = 4 5 6

	7 8 9

Example Output
Output 1:-
12

Output 2:-
45

Example Explanation
Expanation 1:-
The submatrix with max sum is
-1 2 3

	2 2 4
	Sum is 12.

	Explanation 2:-

The largest submatrix with max sum is
1 2 3
4 5 6
7 8 9
The sum is 45.
*/
func GetMaxSubarraySum(A [][]int) int64 {
	var maxSum int64 = math.MinInt64

	//build suffix array starting from from bottom right
	matrixSum := buildSuffixMatrixSum(A)
	for i := range matrixSum {
		for j := range matrixSum[i] {
			if matrixSum[i][j] > maxSum {
				maxSum = matrixSum[i][j]
			}
		}
	}
	return maxSum
}

func buildSuffixMatrixSum(A [][]int) [][]int64 {
	matrixSum := make([][]int64, len(A))
	for i := range matrixSum {
		matrixSum[i] = make([]int64, len(A[0]))
	}

	for i := len(A) - 1; i >= 0; i-- {
		for j := len(A[0]) - 1; j >= 0; j-- {
			matrixSum[i][j] = int64(A[i][j])
			//start from bottom right
			//Formula: matrixSum[i][j] = matrixSum[i][j+1] + matrixSum[i+1][j] - matrixSum[i+1][j+1] + A[i][j]
			if i+j < (len(A) + len(A[0]) - 2) {
				if i == (len(A) - 1) {
					matrixSum[i][j] += matrixSum[i][j+1]
				} else if j == (len(A[0]) - 1) {
					matrixSum[i][j] += matrixSum[i+1][j]
				} else {
					matrixSum[i][j] += matrixSum[i][j+1] + matrixSum[i+1][j] - matrixSum[i+1][j+1]
				}
			}
		}
	}

	return matrixSum
}
