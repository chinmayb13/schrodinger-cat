package arrays1

import "math"

/*
Problem Description
Given 4 array of integers A, B, C and D of same size, find and return the maximum value of | A [ i ] - A [ j ] | + | B [ i ] - B [ j ] | + | C [ i ] - C [ j ] | + | D [ i ] - D [ j ] | + | i - j| where i != j and |x| denotes the absolute value of x.

Problem Constraints
2 <= length of the array A, B, C, D <= 100000
-106 <= A[i] <= 106

Input Format
The arguments given are the integer arrays A, B, C, D.

Output Format
Return the maximum value of | A [ i ] - A [ j ] | + | B [ i ] - B [ j ] | + | C [ i ] - C [ j ] | + | D [ i ] - D [ j ] | + | i - j|

Example Input
Input 1:

	A = [1, 2, 3, 4]
	B = [-1, 4, 5, 6]
	C = [-10, 5, 3, -8]
	D = [-1, -9, -6, -10]

	Input 2:

	A = [1, 2]
	B = [3, 6]
	C = [10, 11]
	D = [1, 6]

Example Output
Output 1:
30

Output 2:
11

Example Explanation
Explanation 1:
Maximum value occurs for i = 0 and j = 1.

Explanation 2:
There is only one possibility for i, j as there are only two elements in the array.
*/
func GetMaxAbsoulteDiffII(A []int, B []int, C []int, D []int) int {
	maxDiff := math.MinInt
	signArr := []int{-1, 1}
	for i := range signArr {
		for j := range signArr {
			for k := range signArr {
				for l := range signArr {
					maxValue := math.MinInt
					minValue := math.MaxInt
					for idx := range A {
						sum := signArr[i]*A[idx] + signArr[j]*B[idx] + signArr[k]*C[idx] + signArr[l]*D[idx] + idx
						if sum > maxValue {
							maxValue = sum
						}
						if sum < minValue {
							minValue = sum
						}
					}
					if (maxValue - minValue) > maxDiff {
						maxDiff = maxValue - minValue
					}
				}
			}
		}
	}
	return maxDiff
}
