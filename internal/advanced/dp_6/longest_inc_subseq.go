package dp6

/*
Problem Description
Find the longest increasing subsequence of a given array of integers, A.

In other words, find a subsequence of array in which the subsequence's elements are in strictly increasing order, and in which the subsequence is as long as possible.

In this case, return the length of the longest increasing subsequence.

Problem Constraints
1 <= length(A) <= 2500
0 <= A[i] <= 2500

Input Format
The first and the only argument is an integer array A.

Output Format
Return an integer representing the length of the longest increasing subsequence.

Example Input
Input 1:
A = [1, 2, 1, 5]

Input 2:
A = [0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15]

Example Output
Output 1:
3

Output 2:
6

Example Explanation
Explanation 1:
The longest increasing subsequence: [1, 2, 5]

Explanation 2:
The possible longest increasing subsequences: [0, 2, 6, 9, 13, 15] or [0, 4, 6, 9, 11, 15] or [0, 4, 6, 9, 13, 15]
*/
func GetLIS(A []int) int {
	inp := A
	N := len(inp)
	memo := make([]int, N+1)
	for i := 1; i <= N; i++ {
		maxVal := 0
		//for ith index, check from 0 to i-1th index if there is an increasing sequence
		for j := 1; j <= i-1; j++ {
			if inp[i-1] > inp[j-1] {
				maxVal = max(maxVal, memo[j])
			}
		}
		memo[i] = 1 + maxVal
	}

	ans := 0
	//search for the max value
	for i := 1; i <= N; i++ {
		ans = max(ans, memo[i])
	}
	return ans
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}
