package binarysearch1

import "sort"

/*
Problem Description
You are given a 2-D matrix C of size A × B.
You need to build a new 1-D array of size A by taking one element from each row of the 2-D matrix in such a way that the cost of the newly built array is minimized.

The cost of an array is the minimum possible value of the absolute difference between any two adjacent elements of the array.
So if the newly built array is X, the element picked from row 1 will become X[1], element picked from row 2 will become X[2], and so on.
Determine the minimum cost of the newly built array.

Problem Constraints
2 <= A <= 1000
2 <= B <= 1000
1 <= C[i][j] <= 106

Input Format
The first argument is an integer A denoting number of rows in the 2-D array.
The second argument is an integer B denoting number of columns in the 2-D array.
The third argument is a 2-D array C of size A x B.

Output Format
Return an integer denoting the minimum cost of the newly build array.

Example Input
Input 1:
A = 2
B = 2
C = [ [8, 4]

	[6, 8] ]

Input 2:
A = 3
B = 2
C = [ [7, 3]

	[2, 1]
	[4, 9] ]

Example Output
Output 1:
0

Output 2:
1

Example Explanation
Explanation 1:
Newly build array : [8, 8]. The minimum cost will be 0 since the absolute difference will be 0(8-8).

Explanation 2:
Newly build array : [3, 2, 4]. The minimum cost will be 1 since the minimum absolute difference will be 1(3 - 2).
*/
func GetMinCost(A int, B int, C [][]int) int {
	var minCost int = 1e6
	for i := range C {
		sort.Ints(C[i])
	}
	for i := 0; i < A-1; i++ {
		for j := range C[i] {
			left := 0
			right := B - 1
			low := left
			high := right
			for low <= high {
				mid := low + (high-low)>>1
				if C[i+1][mid] <= C[i][j] {
					left = mid
					low = mid + 1
				} else {
					right = mid
					high = mid - 1
				}
			}

			minCost = min(minCost, min(getAbsDiff(C[i][j], C[i+1][left]), getAbsDiff(C[i+1][right], C[i][j])))
		}
	}
	return minCost
}

func getAbsDiff(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1 - inp2
	}
	return inp2 - inp1
}
