package sorting3

import "github.com/chinmayb13/schrodinger-cat/internal/helper"

/*
Problem Description
Given an array A of non-negative integers of size N. Find the minimum sub-array Al, Al+1 ,..., Ar such that if we sort(in ascending order) that sub-array, then the whole array should get sorted. If A is already sorted, output -1.

Problem Constraints
1 <= N <= 1000000
1 <= A[i] <= 1000000

Input Format
First and only argument is an array of non-negative integers of size N.

Output Format
Return an array of length two where the first element denotes the starting index(0-based) and the second element denotes the ending index(0-based) of the sub-array. If the array is already sorted, return an array containing only one element i.e. -1.

Example Input
Input 1:
A = [1, 3, 2, 4, 5]

Input 2:
A = [1, 2, 3, 4, 5]

Example Output
Output 1:
[1, 2]

Output 2:
[-1]

Example Explanation
Explanation 1:
If we sort the sub-array A1, A2, then the whole array A gets sorted.

Explanation 2:
A is already sorted.
*/
func GetMaxUnSortedSubArray(A []int) []int {
	maxV := 0
	minV := 1000001
	beg := -1
	end := -1
	for i := 1; i < len(A); i++ {
		if A[i] < helper.Max(A[i-1], maxV) {
			end = i
			minV = helper.Min(minV, A[i])
		}
		maxV = helper.Max(maxV, A[i])
	}

	if end == -1 {
		return []int{-1}
	}

	for i := range A {
		if A[i] > minV {
			beg = i
			break
		}
	}
	return []int{beg, end}
}

func GetMaxUnSortedSubArrayAlt(A []int) []int {
	beg, end := -1, -1
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			beg = i
			break
		}
	}
	if beg == -1 {
		return []int{-1}
	}
	for i := len(A) - 1; i > 0; i-- {
		if A[i] < A[i-1] {
			end = i
			break
		}
	}

	max := 0
	min := 1000001
	for k := beg; k <= end; k++ {
		if A[k] > max {
			max = A[k]
		}
		if A[k] < min {
			min = A[k]
		}
	}

	for k := beg; k >= 0 && min < A[k]; k-- {
		beg = k
	}

	for j := end; j < len(A) && max > A[j]; j++ {
		end = j
	}

	return []int{beg, end}

}
