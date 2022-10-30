package bitmanipulations2

/*
Problem Description
You are given two integers A and B.
Set the A-th bit and B-th bit in 0, and return output in decimal.

Problem Constraints
0 <= A <= 30
0 <= B <= 30

Input Format
First argument A is an integer.
Second argument B is an integer.

Output Format
Return an integer.

Example Input
Input 1:
A = 3
B = 5

Input 2:
A = 4
B = 4

Example Output
Output 1:
40

Output 2:
16

Example Explanation
For Input 1:
The binary expression is 101000 which is 40 in decimal.

For Input 2:
The binary expression is 10000 which is 16 in decimal
*/
func SetTheBits(A int, B int) int {
	return (1 << A) | (1 << B)
}