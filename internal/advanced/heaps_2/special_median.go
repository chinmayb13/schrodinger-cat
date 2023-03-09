package heaps2

import (
	"container/heap"
	"math"
)

/*
Problem Description
You are given an array A containing N numbers. This array is called special if it satisfies one of the following properties:

There exists an element A[i] in the array such that A[i] is equal to the median of elements [A[0], A[1], ...., A[i-1]]
There exists an element A[i] in the array such that A[i] is equal to the median of elements [A[i+1], A[i+2], ...., A[N-1]]
The Median is the middle element in the sorted list of elements. If the number of elements is even then the median will be (sum of both middle elements) / 2.

Return 1 if the array is special else return 0.

NOTE:
Do not neglect decimal point while calculating the median
For A[0] consider only the median of elements [A[1], A[2], ..., A[N-1]] (as there are no elements to the left of it)
For A[N-1] consider only the median of elements [A[0], A[1], ...., A[N-2]]

Problem Constraints
1 <= N <= 1000000.
A[i] is in the range of a signed 32-bit integer.

Input Format
The first and only argument is an integer array A.

Output Format
Return 1 if the given array is special else return 0.

Example Input
Input 1:
A = [4, 6, 8, 4]

Input 2:
A = [2, 7, 3, 1]

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explantion 1:
Here, 6 is equal to the median of [8, 4].

Explanation 2:
No element satisfies any condition.
*/

//idea: do one pass from left and another from right to find the median
func ContainsSpecialMedian(A []int) int {
	if medianFromLeft(A) || medianFromRight(A) {
		return 1
	}
	return 0
}

func medianFromLeft(inp []int) bool {
	N := len(inp)
	mxh := make(maxIntHeap, 0)
	mnh := make(minIntHeap, 0)
	heap.Init(&mxh)
	heap.Init(&mnh)
	median := math.MaxInt
	for i := 0; i < N-1; i++ {
		if inp[i] < median {
			if len(mxh) > len(mnh) {
				heap.Push(&mnh, heap.Pop(&mxh))
			}
			heap.Push(&mxh, inp[i])
		} else {
			if len(mnh) > len(mxh) {
				heap.Push(&mxh, heap.Pop(&mnh))
			}
			heap.Push(&mnh, inp[i])
		}

		if len(mxh) > len(mnh) {
			median = mxh[0]
		} else if len(mnh) > len(mxh) {
			median = mnh[0]
		} else {
			sum := mxh[0] + mnh[0]
			median = sum >> 1
			if sum&1 > 0 {
				continue
			}
		}
		if median == inp[i+1] {
			return true
		}
	}
	return false
}

func medianFromRight(inp []int) bool {
	N := len(inp)
	mxh := make(maxIntHeap, 0)
	mnh := make(minIntHeap, 0)
	heap.Init(&mxh)
	heap.Init(&mnh)
	median := math.MaxInt
	for i := N - 1; i > 0; i-- {
		if inp[i] < median {
			if len(mxh) > len(mnh) {
				heap.Push(&mnh, heap.Pop(&mxh))
			}
			heap.Push(&mxh, inp[i])
		} else {
			if len(mnh) > len(mxh) {
				heap.Push(&mxh, heap.Pop(&mnh))
			}
			heap.Push(&mnh, inp[i])
		}

		if len(mxh) > len(mnh) {
			median = mxh[0]
		} else if len(mnh) > len(mxh) {
			median = mnh[0]
		} else {
			sum := mxh[0] + mnh[0]
			median = sum >> 1
			if sum&1 > 0 {
				continue
			}
		}
		if median == inp[i-1] {
			return true
		}
	}
	return false
}
