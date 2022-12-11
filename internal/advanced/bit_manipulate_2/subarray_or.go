package bitmanipulate2

import "math"

/*
Problem Description
You are given an array of integers A of size N.
The value of a subarray is defined as BITWISE OR of all elements in it.
Return the sum of value of all subarrays of A % 109 + 7.

Problem Constraints
1 <= N <= 105
1 <= A[i] <= 108

Input Format
The first argument given is the integer array A.

Output Format
Return the sum of Value of all subarrays of A % 109 + 7.

Example Input
Input 1:
A = [1, 2, 3, 4, 5]

Input 2:
A = [7, 8, 9, 10]


Example Output
Output 1:
71

Output 2:
110

Example Explanation
Explanation 1:
Value([1]) = 1
Value([1, 2]) = 3
Value([1, 2, 3]) = 3
Value([1, 2, 3, 4]) = 7
Value([1, 2, 3, 4, 5]) = 7
Value([2]) = 2
Value([2, 3]) = 3
Value([2, 3, 4]) = 7
Value([2, 3, 4, 5]) = 7
Value([3]) = 3
Value([3, 4]) = 7
Value([3, 4, 5]) = 7
Value([4]) = 4
Value([4, 5]) = 5
Value([5]) = 5
Sum of all these values = 71

Explanation 2:
Sum of value of all subarray is 110.
*/
/*
Approach: For each bit position, find out how many subarrays have that position unset
contribution = (total number of subarrays - number of subarrays with jth bit unset)*(1<<j)
*/
func GetSubArrayORSum(A []int) int {
	var mod int = 1e9 + 7
	orSum := 0
	max := 0
	for i := range A {
		if A[i] > max {
			max = A[i]
		}
	}
	bits := int(math.Floor(math.Log2(float64(max)) + 1))
	for j := 0; j < bits; j++ {
		zeroSubArrayCount := 0
		zeroIdx := -1
		for i := range A {
			if (A[i] & (1 << j)) == 0 {
				if zeroIdx == -1 {
					zeroIdx = i
				}
				zeroSubArrayCount += i - zeroIdx + 1
			} else {
				zeroIdx = -1
			}
		}
		orSum = (orSum + ((len(A)*(len(A)+1))/2-zeroSubArrayCount)*(1<<j)) % mod

	}
	return orSum
}
