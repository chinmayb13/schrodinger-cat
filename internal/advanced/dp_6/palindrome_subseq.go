package dp6

/*
Problem Description
Given a string A. Find the longest palindromic subsequence (A subsequence which does not need to be contiguous and is a palindrome).
You need to return the length of longest palindromic subsequence.

Problem Constraints
1 <= length of(A) <= 103

Input Format
First and only integer is a string A.

Output Format
Return an integer denoting the length of longest palindromic subsequence.

Example Input
Input 1:
A = "bebeeed"

Input 2:
A = "aedsead"

Example Output
Output 1:
4

Output 2:
5

Example Explanation
Explanation 1:
The longest palindromic subsequence is "eeee", which has a length of 4.

Explanation 2:
The longest palindromic subsequence is "aedea", which has a length of 5.
*/
func GetLongestPalindromeSubSeq(A string) int {
	inp := A
	N := len(inp)
	memo := make([][]int, N+1)
	for i := range memo {
		memo[i] = make([]int, N+1)
	}

	//for every substring initialise the subsequence count to 1
	for i := 1; i <= N; i++ {
		for j := i; j <= N; j++ {
			memo[i][j] = 1
		}
	}

	for i := N - 1; i >= 1; i-- {
		for j := i + 1; j <= N; j++ {
			//take max of right part excluding the ith character
			//and the left part excluding the jth character
			memo[i][j] = max(memo[i][j-1], memo[i+1][j])
			//if i and j are same characters, find if [i+1,j-1] is maximum
			if inp[i-1] == inp[j-1] {
				memo[i][j] = max(memo[i][j], memo[i+1][j-1]+2)
			}
		}

	}

	//return ans for 1st row
	idx := 1
	ans := 0
	for idx <= N {
		ans = max(ans, memo[1][idx])
		idx++
	}
	return ans
}
