package combinatorics

import (
	"strings"
)

/*
Problem Description
Given a positive integer A, return its corresponding column title as it appears in an Excel sheet.

For example:

1 -> A
2 -> B
3 -> C
...
26 -> Z
27 -> AA
28 -> AB

Problem Constraints
1 <= A <= 109

Input Format
First and only argument of input contains single integer A

Output Format
Return a string denoting the corresponding title.

Example Input
Input 1:
A = 3

Input 2:
A = 27

Example Output
Output 1:
"C"

Output 2:
"AA"

Example Explanation
Explanation 1:

3 corrseponds to C.

Explanation 2:
1 -> A,
2 -> B,
3 -> C,
...
26 -> Z,
27 -> AA,
28 -> AB
*/
func GetExcelColumnHeader(A int) string {
	var sb strings.Builder
	for A > 0 {
		A--
		sb.WriteByte(byte((A)%26) + 'A')
		A = A / 26
	}
	return rev([]byte(sb.String()))
}

func rev(inp []byte) string {
	for i, j := 0, len(inp)-1; i < j; i, j = i+1, j-1 {
		inp[i], inp[j] = inp[j], inp[i]
	}
	return string(inp)
}
