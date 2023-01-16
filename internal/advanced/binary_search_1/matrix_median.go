package binarysearch1

/*
Problem Description
Given a matrix of integers A of size N x M in which each row is sorted.
Find and return the overall median of matrix A.
NOTE: No extra memory is allowed.
NOTE: Rows are numbered from top to bottom and columns are numbered from left to right.

Problem Constraints
1 <= N, M <= 10^5
1 <= N*M <= 10^6
1 <= A[i] <= 10^9
N*M is odd

Input Format
The first and only argument given is the integer matrix A.

Output Format
Return the overall median of matrix A.

Example Input
Input 1:
A = [   [1, 3, 5],
        [2, 6, 9],
        [3, 6, 9]   ]

Input 2:
A = [   [5, 17, 100]    ]

Example Output
Output 1:
5

Output 2:
17

Example Explanation
Explanation 1:
A = [1, 2, 3, 3, 5, 6, 6, 9, 9]
Median is 5. So, we return 5.

Explanation 2:
Median is 17.
*/

func FindMedian(A [][]int) int {
	ans := 0
	expectedSmallerCount := (len(A)*len(A[0]) + 1) / 2
	minVal, maxVal := getMinMax(A)
	for minVal <= maxVal {
		mid := minVal + (maxVal-minVal)/2
		smallerCount := getSmallerCount(A, mid)
		if smallerCount >= expectedSmallerCount {
			ans = mid
			maxVal = mid - 1
		} else {
			minVal = mid + 1
		}
	}
	return ans
}

func getMinMax(inp [][]int) (minVal, maxVal int) {
	minVal = 1e9 + 1
	maxVal = 0
	colsize := len(inp[0])
	for i := range inp {
		minVal = min(minVal, inp[i][0])
		maxVal = max(maxVal, inp[i][colsize-1])
	}
	return
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}

func getSmallerCount(inp [][]int, val int) int {
	colsize := len(inp[0])
	ans := 0
	for i := range inp {
		low, high := 0, colsize-1
		idx := -1
		for low <= high {
			mid := low + (high-low)/2
			if inp[i][mid] <= val {
				idx = mid
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		ans += idx + 1
	}
	return ans
}
