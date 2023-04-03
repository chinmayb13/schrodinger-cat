package dp6

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
func GetPalindromeCount(A string) int {
	inp := A
	N := len(inp)
	pMatrix := make([][]bool, N+1)
	for i := range pMatrix {
		pMatrix[i] = make([]bool, N+1)
	}

	//invalid combination to be marked true
	for i := 0; i <= N; i++ {
		for j := 0; j < i; j++ {
			pMatrix[i][j] = true
		}
	}
	//last character of the string to be marked true
	pMatrix[N][N] = true
	ans := 1
	for i := N - 1; i >= 1; i-- {
		for j := i; j <= N; j++ {
			//check if the start and end index match, then check for second and last second index
			pMatrix[i][j] = inp[i-1] == inp[j-1] && pMatrix[i+1][j-1]
			if pMatrix[i][j] {
				ans++
			}
		}
	}
	return ans

}
