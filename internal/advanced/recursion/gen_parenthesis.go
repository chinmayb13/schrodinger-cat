package recursion

import "strings"

/*
Problem Description
Given an integer A pairs of parentheses, write a function to generate all combinations of well-formed parentheses of length 2*A.

Problem Constraints
1 <= A <= 10

Input Format
First and only argument is integer A.

Output Format
Return a sorted list of all possible parenthesis.

Example Input
Input 1:
A = 3

Input 2:
A = 1

Example Output
Output 1:
[ "((()))", "(()())", "(())()", "()(())", "()()()" ]

Output 2:
[ "()" ]


Example Explanation
Explanation 1:
All paranthesis are given in the output list.

Explanation 2:
All paranthesis are given in the output list.
*/
func GenerateParenthesis(A int) []string {
	var ans []string
	var sb strings.Builder
	genParenthesis(A, A, &sb, &ans)
	return ans

}

func genParenthesis(start, end int, sb *strings.Builder, ans *([]string)) {
	if start == 0 && end == 0 {
		*ans = append(*ans, sb.String())
	}
	if start > 0 {
		sb.WriteString("(")
		genParenthesis(start-1, end, sb, ans)
		removeLastChar(sb)
	}

	if end > 0 && start < end {
		sb.WriteString(")")
		genParenthesis(start, end-1, sb, ans)
		removeLastChar(sb)
	}

}

func removeLastChar(sb *strings.Builder) {
	val := sb.String()
	sb.Reset()
	sb.WriteString(val[:len(val)-1])
}
