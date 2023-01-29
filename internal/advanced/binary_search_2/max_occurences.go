package binarysearch2

import "sort"

/*
Problem Description
Given an array of integers A of size N and an integer B.

In a single operation, any one element of the array can be increased by 1. You are allowed to do at most B such operations.

Find the number with the maximum number of occurrences and return an array C of size 2, where C[0] is the number of occurrences, and C[1] is the number with maximum occurrence.
If there are several such numbers, your task is to find the minimum one.

Problem Constraints
1 <= N <= 105
-109 <= A[i] <= 109
0 <= B <= 109

Input Format
The first argument given is the integer array A.
The second argument given is the integer B.

Output Format
Return an array C of size 2, where C[0] is number of occurence and C[1] is the number with maximum occurence.

Example Input
Input 1:
A = [3, 1, 2, 2, 1]
B = 3

Input 2:
A = [5, 5, 5]
B = 3

Example Output
Output 1:
[4, 2]

Output 2:
[3, 5]

Example Explanation
Explanation 1:
Apply operations on A[2] and A[4]
A = [3, 2, 2, 2, 2]
Maximum occurence =  4
Minimum value of element with maximum occurence = 2

Explanation 2:
A = [5, 5, 5]
Maximum occurence =  3
Minimum value of element with maximum occurence = 5
*/

func GetMaxOccurrence(A []int, B int) []int {
	sort.Ints(A)
	ans := []int{-1, -1}
	ps := buildPrefixSum(A, B)
	var low, mid, high int
	//since the element is atleast present once or arr size times at max
	low = 1
	high = len(A)
	for low <= high {
		mid = low + (high-low)>>1
		//check if any element can be repeated window size times
		idx := isWindowPossible(A, ps, B, mid)
		if idx != -1 {
			ans[0] = mid
			ans[1] = A[idx]
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

func buildPrefixSum(inp []int, ops int) []int64 {
	ps := make([]int64, len(inp))
	ps[0] = int64(inp[0])
	for i := 1; i < len(ps); i++ {
		ps[i] = ps[i-1] + int64(inp[i])
	}
	return ps
}

//idea: use sliding window
func isWindowPossible(inp []int, ps []int64, B, window int) int {
	for i := window - 1; i < len(inp); i++ {
		sum := ps[i]
		if i-window >= 0 {
			sum -= ps[i-window]
		}
		//if the last element of the window can be placed in the entire window with the help of B
		if (sum + int64(B)) >= int64(inp[i])*int64(window) {
			return i
		}
	}
	return -1
}
