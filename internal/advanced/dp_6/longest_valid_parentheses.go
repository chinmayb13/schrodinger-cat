package dp6

/*
Problem Description
Given a string A containing just the characters '(' and ')'.

Find the length of the longest valid (well-formed) parentheses substring.

Problem Constraints
1 <= length(A) <= 750000

Input Format
The only argument given is string A.

Output Format
Return the length of the longest valid (well-formed) parentheses substring.

Example Input
Input 1:
A = "(()"

Input 2:
A = ")()())"

Example Output
Output 1:
2

Output 2:
4

Example Explanation
Explanation 1:
The longest valid parentheses substring is "()", which has length = 2.

Explanation 2:
The longest valid parentheses substring is "()()", which has length = 4.
*/
func LongestValidParentheses(A string) int {
	inp := A
	N := len(inp)
	memo := make([]int, N+1)

	for i := 1; i <= N; i++ {
		//check for closing braces
		if i > 1 && inp[i-1] == ')' {
			//if the last character was open brace
			//add 2 to the answer at index before open brace
			if inp[i-2] == '(' {
				memo[i] = memo[i-2] + 2
			}

			//if there was valid parentheses ending at last character
			//check if there is an open brace before the valid parentheses length
			//if it is, save ans before the open brace + valid parentheses length + 2
			if memo[i-1] > 0 && i-memo[i-1]-2 >= 0 && inp[i-memo[i-1]-2] == '(' {
				memo[i] = memo[i-memo[i-1]-2] + memo[i-1] + 2
			}
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans = max(ans, memo[i])
	}
	return ans
}
