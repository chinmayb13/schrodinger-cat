package queue

import "strings"

/*
Problem Description
Given an integer A, you have to find the Ath Perfect Number.
A Perfect Number has the following properties:
It comprises only 1 and 2.
The number of digits in a Perfect number is even.
It is a palindrome number.
For example, 11, 22, 112211 are Perfect numbers, where 123, 121, 782, 1 are not.

Problem Constraints
1 <= A <= 100000

Input Format
The only argument given is an integer A.

Output Format
Return a string that denotes the Ath Perfect Number.

Example Input
Input 1:
A = 2

Input 2:
A = 3

Example Output
Output 1:
22

Output 2:
1111

Example Explanation
Explanation 1:
First four perfect numbers are:
1. 11
2. 22
3. 1111
4. 1221

Explanation 2:
First four perfect numbers are:
1. 11
2. 22
3. 1111
4. 1221
*/
func GetAthPerfectNumber(A int) string {
	perfectNumbers := make([]string, A+1)
	perfectNumbers[0] = "1"
	perfectNumbers[1] = "2"
	front, back := 0, 1
	var sb strings.Builder
	for back < A-1 {
		source := perfectNumbers[front]
		sb.WriteString(source)
		sb.WriteString("1")
		perfectNumbers[back+1] = sb.String()
		sb.Reset()
		sb.WriteString(source)
		sb.WriteString("2")
		perfectNumbers[back+2] = sb.String()
		sb.Reset()
		front++
		back += 2
	}
	sb.WriteString(perfectNumbers[A-1])
	sb.WriteString(reverseString(perfectNumbers[A-1]))
	return sb.String()
}

func reverseString(s string) string {
	var rev strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		rev.WriteByte(s[i])
	}
	return rev.String()
}
