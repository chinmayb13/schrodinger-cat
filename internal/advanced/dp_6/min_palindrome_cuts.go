package dp6

import "math"

/*
Given a string A consisting of lowercase English alphabets. Your task is to find how many substrings of A are palindrome.
The substrings with different start indexes or end indexes are counted as different substrings even if they consist of same characters.
Return the count of palindromic substrings.

Note: A string is palindrome if it reads the same from backward and forward.

Input Format
The only argument given is string A.
Output Format

Return the count of palindromic substrings.
Constraints

1 <= length of the array <= 1000
For Example
Input 1:
A = "abab"

Output 1:
6

Explanation 1:
6 palindromic substrings are:
"a", "aba", "b", "bab", "a" and "b".

Input 2:
A = "ababa"

Output 2:
9

Explanation 9:
9 palindromic substrings are:
"a", "a", "a", "b", "b" , "aba" ,"bab", "aba" and "ababa".
*/
func MinCut(A string) int {
	//create a substring palindrome matrix
	pMatrix := getPalidromeMatrix(A)
	N := len(A)
	memo := make([]int, N)

	for i := 0; i < N; i++ {
		//if the substring itself is a palindrome
		if pMatrix[1][i+1] {
			continue
		}
		ans := math.MaxInt
		for j := 0; j < i; j++ {
			//check for substring [0,j] only when [j+1,i] is a valid palindrome
			//take min of them
			if pMatrix[j+2][i+1] {
				ans = min(ans, memo[j]+1)
			}
		}
		memo[i] = ans
	}
	//return the min cuts for the whole string
	return memo[N-1]
}

func getPalidromeMatrix(inp string) [][]bool {
	N := len(inp)
	pMatrix := make([][]bool, N+1)
	for i := range pMatrix {
		pMatrix[i] = make([]bool, N+1)
	}

	for i := 0; i <= N; i++ {
		for j := 0; j < i; j++ {
			pMatrix[i][j] = true
		}
	}
	pMatrix[N][N] = true
	for i := N - 1; i >= 1; i-- {
		for j := i; j <= N; j++ {
			pMatrix[i][j] = inp[i-1] == inp[j-1] && pMatrix[i+1][j-1]
		}
	}
	return pMatrix
}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}
