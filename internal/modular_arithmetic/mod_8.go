package modulararithmetic

import (
	"strconv"
	"strings"
)

/*
Problem Description
You are given a number A in the form of a string. Check if the number is divisible by eight or not.

Return 1 if it is divisible by eight else, return 0.

Problem Constraints
1 <= length of the String <= 100000
'0' <= A[i] <= '9'

Input Format
The only argument given is a string A.

Output Format
Return 1 if it is divisible by eight else return 0.

Example Input
Input 1:
A = "16"

Input 2:
A = "123"

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
16 = 8 * 2

Explanation 2:
123 = 15 * 8 + 3
*/

func IsDivisbleBy8(A string) int {
	var threeChars strings.Builder
	threeChars.Grow(len(A))
	var byteArr []byte
	//last 3 digits
	for i, j := len(A)-1, 1; i >= 0 && j <= 3; i, j = i-1, j+1 {
		byteArr = append(byteArr, A[i])
	}
	//reverse to get the actual digit order "61" to "16"
	for i := len(byteArr) - 1; i >= 0; i-- {
		threeChars.WriteByte(byteArr[i])
	}

	threeDigit, _ := strconv.Atoi(threeChars.String())
	if threeDigit%8 == 0 {
		return 1
	}
	return 0
}
