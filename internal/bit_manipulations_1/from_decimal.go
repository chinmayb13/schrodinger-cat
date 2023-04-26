package bitmanipulations1

/*
Problem Description
Given a decimal number A and base B.
You are required to change the decimal number A into the corresponding value in base B and return that.


Problem Constraints
0 <= A <= 512
2 <= B <= 10


Input Format
The first argument will be decimal number A.
The second argument will be base B.


Output Format
Return the conversion of A in base B.


Example Input
Input:
A = 4
B = 3


Example Output
Output:
11


Example Explanation
Explanation:
Decimal number 4 in base 3 is 11.
*/

// Approach: Keep on dividing the number by base and multiply the remainder by increasing power of 10 and store it into a sum
func DecimalToAnyBase(A int, B int) int {
	quotient := A
	number := 0
	exponent := 1
	for {
		if quotient == 0 {
			break
		}
		remainder := quotient % B
		number += remainder * exponent
		quotient = quotient / B
		exponent *= 10
	}
	return number
}
