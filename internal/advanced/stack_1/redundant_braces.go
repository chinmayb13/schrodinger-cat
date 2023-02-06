package stack1

/*
Problem Description
Given a string A denoting an expression. It contains the following operators '+', '-', '*', '/'.
Check whether A has redundant braces or not.
NOTE: A will be always a valid expression and will not contain any white spaces.

Problem Constraints
1 <= |A| <= 105

Input Format
The only argument given is string A.

Output Format
Return 1 if A has redundant braces else, return 0.

Example Input
Input 1:
A = "((a+b))"

Input 2:
A = "(a+(a+b))"

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
((a+b)) has redundant braces so answer will be 1.

Explanation 2:
(a+(a+b)) doesn't have have any redundant braces so answer will be 0.
*/
func ContainsRedundantBraces(A string) int {
	var count int
	exprStack := newStack()
	for i := range A {
		// if the parentheses has an expression
		if A[i] == ')' {
			for exprStack.topByte() != '(' {
				count++
				exprStack.pop()
			}
			//if no operation was involved inside parentheses
			if count <= 1 {
				return 1
			}
			exprStack.pop()
			//push till ')' is encountered
		} else {
			exprStack.push(A[i])
		}
	}
	return 0
}
