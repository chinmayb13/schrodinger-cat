package dp5

import "math"

/*
Problem Description
Given two strings A and B. Find the longest common subsequence ( A sequence which does not need to be contiguous), which is common in both the strings.

You need to return the length of such longest common subsequence.

Problem Constraints
1 <= Length of A, B <= 1005

Input Format
First argument is a string A.
Second argument is a string B.

Output Format
Return an integer denoting the length of the longest common subsequence.

Example Input
Input 1:
A = "abbcdgf"
B = "bbadcgf"

Input 2:
A = "aaaaaa"
B = "ababab"

Example Output
Output 1:
5

Output 2:
3

Example Explanation
Explanation 1:
The longest common subsequence is "bbcgf", which has a length of 5.

Explanation 2:
The longest common subsequence is "aaa", which has a length of 3.
*/
func GetLCS(A string, B string) int {
	source := A
	target := B
	N := len(source)
	M := len(target)

	memo := make([][]int, N+1)
	for i := range memo {
		memo[i] = make([]int, M+1)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if source[i-1] == target[j-1] {
				memo[i][j] = 1 + memo[i-1][j-1]
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i][j-1])
			}
		}
	}
	return memo[N][M]
}

func max(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	maxVal := math.MinInt
	for i := range args {
		if args[i] > maxVal {
			maxVal = args[i]
		}
	}
	return maxVal
}
