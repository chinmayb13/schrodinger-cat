package stringspatternmatching

import "strings"

/*
Problem Description
You are given a string A of size N.
Return the string A after reversing the string word by word.

NOTE:
A sequence of non-space characters constitutes a word.
Your reversed string should not contain leading or trailing spaces, even if it is present in the input string.
If there are multiple spaces between words, reduce them to a single space in the reversed string.

Problem Constraints
1 <= N <= 3 * 105

Input Format
The only argument given is string A.

Output Format
Return the string A after reversing the string word by word.

Example Input
Input 1:
A = "the sky is blue"

Input 2:
A = "this is ib"

Example Output
Output 1:
"blue is sky the"

Output 2:
"ib is this"

Example Explanation
Explanation 1:
We reverse the string word by word so the string becomes "blue is sky the".

Explanation 2:
We reverse the string word by word so the string becomes "ib is this".
*/
func GetReverseString(A string) string {
	var sb strings.Builder
	beg, end := findNonSpaceIdx(A)
	var prev int
	cur := end
	for cur >= beg {
		//find the end position of the word
		for cur >= 0 && A[cur] == byte(' ') {
			cur--
		}

		//set the prev to end position + 1
		prev = cur + 1

		//find the start position of the word
		for cur >= 0 && A[cur] != byte(' ') {
			cur--
		}

		sb.WriteString(A[cur+1 : prev])
		if cur >= beg {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func findNonSpaceIdx(inp string) (int, int) {
	beg := 0
	end := len(inp) - 1
	for beg < len(inp)-1 && inp[beg] == ' ' {
		beg++
	}

	for end >= 0 && inp[end] == ' ' {
		end--
	}
	return beg, end
}
