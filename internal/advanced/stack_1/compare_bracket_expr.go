package stack1

import "reflect"

/*
Problem Description
Given two strings A and B. Each string represents an expression consisting of lowercase English alphabets, '+', '-', '(' and ')'.
The task is to compare them and check if they are similar. If they are identical, return 1 else, return 0.
NOTE: It may be assumed that there are at most 26 operands from ‘a’ to ‘z’, and every operand appears only once.

Problem Constraints
1 <= length of the each String <= 100

Input Format
The given arguments are string A and string B.

Output Format
Return 1 if they represent the same expression else return 0.

Example Input
Input 1:
A = "-(a+b+c)"
B = "-a-b-c"

Input 2:
A = "a-b-(c-d)"
B = "a-b-c-d"

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
The expression "-(a+b+c)" can be written as "-a-b-c" which is equal as B.

Explanation 2:
Both the expression are different.
*/

func AreExprSame(A string, B string) int {
	if reflect.DeepEqual(evalExp(A), evalExp(B)) {
		return 1
	}
	return 0
}

func evalExp(inp string) []bool {
	//bool array to store sign for every character
	signArr := make([]bool, 26)
	//to store the how many minuses have been encountered
	var minusCount int
	//to store the number of open parantheses
	negSign := newStack()
	for j := 0; j < len(inp); j++ {
		if inp[j] == '(' {
			//if '-(' is encountered, increment minusCount and and push true parentheses stack
			if j > 0 && inp[j-1] == '-' {
				minusCount++
				negSign.push(true)
				//else if '(' or '+(', push false parentheses stack,
			} else {
				negSign.push(false)
			}
			//if ')' is encountered, check the corresponding starting parentheses was negative or not
			//if negative, decrement the minusCount
		} else if inp[j] == ')' {
			if negSign.topBool() {
				minusCount--
			}
			negSign.pop()
		} else if inp[j] >= 'a' && inp[j] <= 'z' {
			//if '-<char> is encountered, increase the minusCount'
			if j > 0 && inp[j-1] == '-' {
				minusCount++
				//check if the minusCount is odd or even
				signArr[int(inp[j]-'a')] = minusCount&1 == 0
				minusCount--
			} else {
				//check if the minusCount is odd or even
				signArr[int(inp[j]-'a')] = minusCount&1 == 0
			}
		}
	}

	return signArr
}
