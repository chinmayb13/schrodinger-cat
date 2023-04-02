package dp5

import "math"

/*
Problem Description
Given two strings A and B, find the minimum number of steps required to convert A to B. (each operation is counted as 1 step.)

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character

Problem Constraints
1 <= length(A), length(B) <= 450

Input Format
The first argument of input contains a string, A.
The second argument of input contains a string, B.

Output Format
Return an integer, representing the minimum number of steps required.

Example Input
Input 1:
A = "abad"
B = "abac"

Input 2:
A = "Anshuman"
B = "Antihuman

Example Output
Output 1:
1

Output 2:
2

Example Explanation
Exlanation 1:
A = "abad" and B = "abac"
After applying operation: Replace d with c. We get A = B.

Explanation 2:
A = "Anshuman" and B = "Antihuman"
After applying operations: Replace s with t and insert i before h. We get A = B.
*/

func GetMinDistance(A string, B string) int {
	source := A
	target := B
	N := len(source)
	M := len(target)

	memo := make([][]int, N+1)
	for i := range memo {
		memo[i] = make([]int, M+1)
		//if the elements in the second string are exhausted,
		//take whatever length is of the first string remaining
		memo[i][0] = i
	}

	memo[0][0] = 0

	//if the elements in the first string are exhausted,
	//take whatever length is of the second string remaining
	for i := 1; i <= M; i++ {
		memo[0][i] = i
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			//if the characters match
			if source[i-1] == target[j-1] {
				memo[i][j] = memo[i-1][j-1]
				//else take min of replace, delete and insert
			} else {
				memo[i][j] = 1 + min(memo[i-1][j-1], memo[i-1][j], memo[i][j-1])
			}
		}
	}
	return memo[N][M]
}

func min(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	minVal := math.MaxInt
	for i := range args {
		if args[i] < minVal {
			minVal = args[i]
		}
	}
	return minVal
}
