package twopointers

import "github.com/chinmayb13/schrodinger-cat/internal/helper"

/*
Problem Description
You are given 3 sorted arrays A, B and C.
Find i, j, k such that : max(abs(A[i] - B[j]), abs(B[j] - C[k]), abs(C[k] - A[i])) is minimized.
Return the minimum max(abs(A[i] - B[j]), abs(B[j] - C[k]), abs(C[k] - A[i])).

Problem Constraints
0 <= len(A), len(B), len(c) <= 106
0 <= A[i], B[i], C[i] <= 107

Input Format
First argument is an integer array A.
Second argument is an integer array B.
Third argument is an integer array C.

Output Format
Return an single integer denoting the minimum max(abs(A[i] - B[j]), abs(B[j] - C[k]), abs(C[k] - A[i])).

Example Input
Input 1:
A = [1, 4, 10]
B = [2, 15, 20]
C = [10, 12]

Input 2:
A = [3, 5, 6]
B = [2]
C = [3, 4]

Example Output
Output 1:
5

Output 2:
1

Example Explanation
Explanation 1:
With 10 from A, 15 from B and 10 from C.

Explanation 2:
With 3 from A, 2 from B and 3 from C.
*/
func GetMinMaxAbsDiff(A []int, B []int, C []int) int {
	var minSoFar int = 1e6 + 1
	i, j, k := 0, 0, 0
	sizeA, sizeB, sizeC := len(A), len(B), len(C)
	for i < sizeA && j < sizeB && k < sizeC {
		localMin := helper.Min(C[k], helper.Min(A[i], B[j]))
		localMax := helper.Max(C[k], helper.Max(A[i], B[j]))
		minSoFar = helper.Min(minSoFar, localMax-localMin)
		if localMin == A[i] {
			i++
		}
		if localMin == B[j] {
			j++
		}
		if localMin == C[k] {
			k++
		}
	}
	if minSoFar == 1e6+1 {
		return 0
	}
	return minSoFar

}
