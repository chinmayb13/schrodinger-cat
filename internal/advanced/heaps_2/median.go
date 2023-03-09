package heaps2

import (
	"container/heap"
	"math"
)

/*
Problem Description
Given an array of integers, A denoting a stream of integers. New arrays of integer B and C are formed.
Each time an integer is encountered in a stream, append it at the end of B and append the median of array B at the C.
Find and return the array C.

NOTE:
If the number of elements is N in B and N is odd, then consider the median as B[N/2] ( B must be in sorted order).
If the number of elements is N in B and N is even, then consider the median as B[N/2-1]. ( B must be in sorted order).

Problem Constraints
1 <= length of the array <= 100000
1 <= A[i] <= 109

Input Format
The only argument given is the integer array A.

Output Format
Return an integer array C, C[i] denotes the median of the first i elements.

Example Input
Input 1:
A = [1, 2, 5, 4, 3]

Input 2:
A = [5, 17, 100, 11]

Example Output
Output 1:
[1, 1, 2, 2, 3]

Output 2:
[5, 5, 17, 11]

Example Explanation
Explanation 1:
stream          median
[1]             1
[1, 2]          1
[1, 2, 5]       2
[1, 2, 5, 4]    2
[1, 2, 5, 4, 3] 3

Explanation 2:
stream          median
[5]              5
[5, 17]          5
[5, 17, 100]     17
[5, 17, 100, 11] 11
*/
func GetMedian(A []int) []int {
	median := math.MaxInt
	var ans []int
	mxh := make(maxIntHeap, 0)
	mnh := make(minIntHeap, 0)
	heap.Init(&mnh)
	heap.Init(&mxh)
	for i := range A {
		//if the new element is trying to go to left bucket
		if A[i] < median {
			//if the left bucket size is more than the right bucket, shift the max from left to right
			if len(mxh) > len(mnh) {
				heap.Push(&mnh, heap.Pop(&mxh))
			}
			//add the new element in the left bucket
			heap.Push(&mxh, A[i])
			//if the new element is trying to go to right bucket
		} else {
			//if the right bucket size is more than the left bucket, shift the min from right to left
			if len(mnh) > len(mxh) {
				heap.Push(&mxh, heap.Pop(&mnh))
			}
			//add the new element in the right bucket
			heap.Push(&mnh, A[i])
		}
		//choose median and save ans
		if len(mnh) > len(mxh) {
			median = mnh[0]
			ans = append(ans, mnh[0])
		} else {
			median = mxh[0]
			ans = append(ans, mxh[0])
		}
	}
	return ans
}
