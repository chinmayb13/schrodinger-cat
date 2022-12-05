package arrays1

/*
Problem Description
You are given an array of N integers, A1, A2, .... AN.
Return the maximum value of f(i, j) for all 1 ≤ i, j ≤ N. f(i, j) is defined as |A[i] - A[j]| + |i - j|, where |x| denotes absolute value of x.

Problem Constraints
1 <= N <= 100000
-109 <= A[i] <= 109

Input Format
First argument is an integer array A of size N.

Output Format
Return an integer denoting the maximum value of f(i, j).

Example Input
Input 1:
A = [1, 3, -1]

Input 2:
A = [2]

Example Output
Output 1:
5

Output 2:
0

Example Explanation
Explanation 1:
f(1, 1) = f(2, 2) = f(3, 3) = 0
f(1, 2) = f(2, 1) = |1 - 3| + |1 - 2| = 3
f(1, 3) = f(3, 1) = |1 - (-1)| + |1 - 3| = 4
f(2, 3) = f(3, 2) = |3 - (-1)| + |2 - 3| = 5
So, we return 5.

Explanation 2:
Only possibility is i = 1 and j = 1. That gives answer = 0.
*/
func GetMaxAbsoulteDiff(A []int) int {
	var xMax int = -(1e9 + 1)
	var xMin int = (1e9 + 1)
	var yMax int = -(1e9 + 1)
	var yMin int = (1e9 + 1)
	for i := range A {
		x := A[i] + i
		y := A[i] - i
		if x > xMax {
			xMax = x
		}
		if x < xMin {
			xMin = x
		}

		if y > yMax {
			yMax = y
		}
		if y < yMin {
			yMin = y
		}

	}
	if (yMax - yMin) > (xMax - xMin) {
		return yMax - yMin
	}
	return xMax - xMin
}
