package bitmanipulate1

import "math"

/*
Problem Description
Divide two integers without using multiplication, division and mod operator.
Return the floor of the result of the division.
Also, consider if there can be overflow cases i.e output is greater than INT_MAX, return INT_MAX.
NOTE: INT_MAX = 2^31 - 1

Problem Constraints
-231 <= A, B <= 231-1
B != 0

Input Format
The first argument is an integer A denoting the dividend.
The second argument is an integer B denoting the divisor.

Output Format
Return an integer denoting the floor value of the division.

Example Input
Input 1:
A = 5
B = 2

Input 2:
A = 7
B = 1

Example Output
Output 1:
2

Output 2:
7

Example Explanation
Explanation 1:
floor(5/2) = 2
*/
func Divide(A int, B int) int {
	var quotient int64 = 0
	var sign int64 = 1
	if A < 0 {
		sign *= -1
		A *= -1
	}
	if B < 0 {
		sign *= -1
		B *= -1
	}
	//start moving from MSB to LSB
	for i := 31; i >= 0; i-- {
		//if the multiple is smaller than A
		if (B << i) <= A {
			//get the quotient
			quotient += 1 << i
			A -= (B << i)
		}
	}
	quotient *= sign
	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(quotient)
}
