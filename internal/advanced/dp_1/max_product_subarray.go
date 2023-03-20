package dp1

import "math"

/*
Problem Description
Given an integer array A of size N. Find the contiguous subarray within the given array (containing at least one number) which has the largest product.
Return an integer corresponding to the maximum product possible.

NOTE: Answer will fit in 32-bit integer value.

Problem Constraints
1 <= N <= 5 * 105
-100 <= A[i] <= 100

Input Format
First and only argument is an integer array A.

Output Format
Return an integer corresponding to the maximum product possible.

Example Input
Input 1:
A = [4, 2, -5, 1]

Input 2:
A = [-3, 0, -5, 0]

Example Output
Output 1:
8

Output 2:
0

Example Explanation
Explanation 1:
We can choose the subarray [4, 2] such that the maximum product is 8.

Explanation 2:
0 will be the maximum product possible.
*/
func MaxProduct(A []int) int {
	ans := math.MinInt
	minP, maxP := 1, 1
	for i := 0; i < len(A); i++ {
		//if current element is negative, min and max to be swapped
		if A[i] < 0 {
			minP, maxP = maxP, minP
		}
		localMinP := minP * A[i]
		localMaxP := maxP * A[i]
		maxP = maxInt(A[i], localMaxP)
		minP = minInt(A[i], localMinP)

		if maxP > ans {
			ans = maxP
		}
	}
	return ans
}

func maxInt(args ...int) int {
	max := math.MinInt
	for i := range args {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}

func minInt(args ...int) int {
	min := math.MaxInt
	for i := range args {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}
