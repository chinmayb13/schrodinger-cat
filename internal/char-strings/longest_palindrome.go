package charstrings

/*
Problem Description
Given a string A of size N, find and return the longest palindromic substring in A.
Substring of string A is A[i...j] where 0 <= i <= j < len(A)

Palindrome string:
A string which reads the same backwards. More formally, A is palindrome if reverse(A) = A.
Incase of conflict, return the substring which occurs first ( with the least starting index).

Problem Constraints
1 <= N <= 6000

Input Format
First and only argument is a string A.

Output Format
Return a string denoting the longest palindromic substring of string A.

Example Input
A = "aaaabaaa"

Example Output
"aaabaaa"

Example Explanation
We can see that longest palindromic substring is of length 7 and the string is "aaabaaa".
*/
func LongestPalindrome(A string) string {
	var start, end int
	for i := 0; i < len(A); i++ {
		l, r := expand(i-1, i+1, A)
		if (r - l) > (end - start) {
			start = l
			end = r
		}

		if i < (len(A) - 1) {
			l, r := expand(i, i+1, A)
			if (r - l) > (end - start) {
				start = l
				end = r
			}
		}
	}
	return A[start:end]
}

func expand(l, r int, A string) (int, int) {
	for l >= 0 && r < len(A) {
		if A[l] != A[r] {
			break
		}
		l--
		r++
	}
	return l + 1, r
}
