package binarysearch1

import "github.com/chinmayb13/schrodinger-cat/internal/helper"

/*
Problem Description
Given an array of integers A, find and return the peak element in it. An array element is peak if it is NOT smaller than its neighbors.
For corner elements, we need to consider only one neighbor. We ensure that answer will be unique.
NOTE: Users are expected to solve this in O(log(N)) time. The array may have duplicate elements.

Problem Constraints
1 <= |A| <= 100000
1 <= A[i] <= 109

Input Format
The only argument given is the integer array A.

Output Format
Return the peak element.

Example Input
Input 1:
A = [1, 2, 3, 4, 5]

Input 2:
A = [5, 17, 100, 11]

Example Output
Output 1:
5

Output 2:
100

Example Explanation
Explanation 1:
5 is the peak.

Explanation 2:
100 is the peak.
*/
func GetPeakElement(A []int) int {
	low, high := 0, len(A)-1
	var mid int
	for low <= high {
		var left, right int
		mid = low + (high-low)/2
		if mid > 0 {
			left = A[mid-1]
		}
		if mid < (len(A) - 1) {
			right = A[mid+1]
		}
		if A[mid] == helper.Max(A[mid], helper.Max(left, right)) {
			return A[mid]
		}
		if left >= right {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
