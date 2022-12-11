package bitmanipulate2

import "math"

/*
Problem Description
Given an array A. For every pair of indices i and j (i != j), find the maximum A[i] & A[j].

Problem Constraints
1 <= len(A) <= 105
1 <= A[i] <= 109

Input Format
The first argument is an integer array A.

Output Format
Return a single integer that is the maximum A[i] & A[j].

Example Input
Input 1:-
A = [53, 39, 88]

Input 2:-
A = [38, 44, 84, 12]

Example Output
Output 1:-
37

Output 2:-
36

Example Explanation
Explanation 1:-
53 & 39 = 37
39 & 88 = 0
53 & 88 = 16
Maximum among all these pairs is 37

Explanation 2:-
Maximum bitwise and among all pairs is (38, 44) = 36
*/
//approach is to find the elements which have the common bit set from MSB end
func GetMaxPairAnd(A []int) int {
	ans := 0
	max := math.MinInt32
	for i := range A {
		if A[i] > max {
			max = A[i]
		}
	}
	bits := int(math.Floor(math.Log2(float64(max)) + 1))
	for j := bits - 1; j >= 0; j-- {
		setBits := 0
		for i := range A {
			setBits += ((A[i] >> j) & 1)
		}
		if setBits > 1 {
			ans |= (1 << j)
			for i := range A {
				if ((A[i] >> j) & 1) == 0 {
					A[i] = 0
				}
			}
		}
	}
	return ans
}
