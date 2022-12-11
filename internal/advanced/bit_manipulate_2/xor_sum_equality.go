package bitmanipulate2

import "math"

/*
Problem Description
Given an integer A.
Two numbers, X and Y, are defined as follows:

X is the greatest number smaller than A such that the XOR sum of X and A is the same as the sum of X and A.
Y is the smallest number greater than A, such that the XOR sum of Y and A is the same as the sum of Y and A.
Find and return the XOR of X and Y.

NOTE 1: XOR of X and Y is defined as X ^ Y where '^' is the BITWISE XOR operator.

NOTE 2: Your code will be run against a maximum of 100000 Test Cases.
*/

/*
Approach:
smaller no greater than A would be would have MSB positioned to the adjacent left of MSB of A, with all other bits as 0
greatest no smaller than A would be the toggled form of A with MSB positioned to the adjacent right of MSB of A
E.g:
5: 101
smaller no greater than A: 1000
greatest no smaller than A: 010
xor: 1010
*/
func GetXorSum(A int) int {
	//calcula
	bits := int(math.Floor(math.Log2(float64(A)) + 1))
	xor := 1 << bits
	for i := 0; (1 << i) <= A; i++ {
		if A&(1<<i) == 0 {
			xor |= (1 << i)
		}
	}
	return xor
}
