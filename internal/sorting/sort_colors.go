package sorting

import "sort"

/*
Problem Description
Given an array with N objects colored red, white, or blue, sort them so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
We will use the integers 0, 1, and 2 to represent red, white, and blue, respectively.
Note: Using the library sort function is not allowed.

Problem Constraints
1 <= N <= 1000000
0 <= A[i] <= 2

Input Format
First and only argument of input contains an integer array A.

Output Format
Return an integer array in asked order

Example Input
Input 1 :
A = [0 1 2 0 1 2]

Input 2:
A = [0]

Example Output
Output 1:
[0 0 1 1 2 2]

Output 2:
[0]

Example Explanation
Explanation 1:
[0 0 1 1 2 2] is the required order.
*/
func GetSortedColorsAlt(A []int) []int {
	sort.Slice(A, func(i, j int) bool {
		return (A[i] == 0 && A[j] != 0) || (A[i] == 1 && A[j] == 2)
	})
	return A
}

func GetSortedColors(A []int) []int {
	count0, count1, count2 := 0, 0, 0
	for i := range A {
		switch A[i] {
		case 0:
			count0++
		case 1:
			count1++
		default:
			count2++
		}
	}

	for i := range A {
		switch {
		case count0 > 0:
			A[i] = 0
			count0--
		case count1 > 0:
			A[i] = 1
			count1--
		default:
			A[i] = 2
		}
	}

	return A
}
