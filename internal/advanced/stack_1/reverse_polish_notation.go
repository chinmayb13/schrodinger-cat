package stack1

import "fmt"

/*
Problem Description
An arithmetic expression is given by a string array A of size N. Evaluate the value of an arithmetic expression in Reverse Polish Notation.
Valid operators are +, -, *, /. Each string may be an integer or an operator.

Problem Constraints
1 <= N <= 105

Input Format
The only argument given is string array A.

Output Format
Return the value of arithmetic expression formed using reverse Polish Notation.

Example Input
Input 1:
    A =   ["2", "1", "+", "3", "*"]

Input 2:
    A = ["4", "13", "5", "/", "+"]

Example Output
Output 1:
9

Output 2:
6

Example Explanation
Explaination 1:
starting from backside:
* : () * ()
3 : () * (3)
+ : (() + ()) * (3)
1 : (() + (1)) * (3)
2 : ((2) + (1)) * (3)
((2) + (1)) * (3) = 9

Explaination 2:
+ : () + ()
/ : () + (() / ())
5 : () + (() / (5))
13 : () + ((13) / (5))
4 : (4) + ((13) / (5))
(4) + ((13) / (5)) = 6
*/
func EvalRPN(A []string) int {
	exprStack := newStack()
	operMap := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	for i := range A {
		if operMap[A[i]] {
			operandA := exprStack.topStrInt()
			exprStack.pop()
			operandB := exprStack.topStrInt()
			exprStack.pop()

			var res int
			switch A[i] {
			case "+":
				res = operandB + operandA
			case "-":
				res = operandB - operandA
			case "*":
				res = operandB * operandA
			case "/":
				res = operandB / operandA
			}
			exprStack.push(fmt.Sprintf("%d", res))
		} else {
			exprStack.push(A[i])
		}
	}
	return exprStack.topStrInt()
}
