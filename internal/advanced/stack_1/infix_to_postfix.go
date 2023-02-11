package stack1

import "strings"

// Problem Description
// Given string A denoting an infix expression. Convert the infix expression into a postfix expression.
// String A consists of ^, /, *, +, -, (, ) and lowercase English alphabets where lowercase English alphabets are operands and ^, /, *, +, - are operators.

// Find and return the postfix expression of A.

// NOTE:
// ^ has the highest precedence.
// / and * have equal precedence but greater than + and -.
// + and - have equal precedence and lowest precedence among given operators.

// Problem Constraints
// 1 <= length of the string <= 500000

// Input Format
// The only argument given is string A.

// Output Format
// Return a string denoting the postfix conversion of A.

// Example Input
// Input 1:
// A = "x^y/(a*z)+b"

// Input 2:
// A = "a+b*(c^d-e)^(f+g*h)-i"

// Example Output
// Output 1:
// "xy^az*/b+"

// Output 2:
// "abcd^e-fgh*+^*+i-"

// Example Explanation
// Explanation 1:
// Output dentotes the postfix expression of the given input.

func GetPostFix(A string) string {
	pMap := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'^': 3,
	}
	var sb strings.Builder
	stack := newStack()
	for i := 0; i < len(A); i++ {
		switch A[i] {
		case '(':
			stack.push(A[i])
		case ')':
			for stack.size() > 0 && stack.topByte() != '(' {
				sb.WriteByte(stack.topByte())
				stack.pop()
			}
			stack.pop()
		case '+', '-', '*', '/', '^':
			//pop the operators which have higher or equal precedence
			for stack.size() > 0 && pMap[A[i]] <= pMap[stack.topByte()] {
				sb.WriteByte(stack.topByte())
				stack.pop()
			}
			stack.push(A[i])
		default:
			sb.WriteByte(A[i])
		}
	}
	for stack.size() > 0 {
		sb.WriteByte(stack.topByte())
		stack.pop()
	}
	return sb.String()
}
