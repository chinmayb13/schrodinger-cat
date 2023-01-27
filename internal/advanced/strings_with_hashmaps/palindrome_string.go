package stringswithhashmaps

import "strings"

/*
Problem Description
Given a string A of size N consisting only of lowercase alphabets. The only operation allowed is to insert characters in the beginning of the string.
Find and return how many minimum characters are needed to be inserted to make the string a palindrome string.

Problem Constraints
1 <= N <= 106

Input Format
The only argument given is a string A.

Output Format
Return an integer denoting the minimum characters needed to be inserted in the beginning to make the string a palindrome string.

Example Input
Input 1:
A = "abc"

Input 2:
A = "bb"

Example Output
Output 1:
2

Output 2:
0

Example Explanation
Explanation 1:
Insert 'b' at beginning, string becomes: "babc".
Insert 'c' at beginning, string becomes: "cbabc".

Explanation 2:
There is no need to insert any character at the beginning as the string is already a palindrome.
*/

func MinPalindromeChars(A string) int {
	var sb strings.Builder
	sb.WriteString(A)
	sb.WriteByte('@')
	for i := len(A) - 1; i >= 0; i-- {
		sb.WriteByte(A[i])
	}

	/*
		String: edec
		Reverse: cede
		input for lps: edec@cede
		LPS: 001000123
		The last value 3 means that first 3 values are palindrome
		In order to make the full string palindrome, number of values to prepend = len(string) - 3
	*/
	lps := computeLPSArray(sb.String())
	return len(A) - lps[len(lps)-1]
}

func MinPalindromeCharsAlt(A string) int {
	ans := 0
	arrLength := len(A)
	l, r := 0, arrLength-1
	for l <= r {
		if A[l] != A[r] {
			//if l==0, there is nothing we could do
			//count the characters from end till current element
			if l == 0 {
				ans = arrLength - r
				r--
				//if l!=0, there is chance that current element can match with l==0
				//since the current element has neither been discarded nor been accepted
				//count the characters from end till element before current element
			} else {
				ans = arrLength - r - 1
				l = 0
			}
			//palindrome case
		} else {
			l++
			r--
		}
	}
	return ans
}
