package heaps2

import (
	"container/heap"
	"sort"
)

/*
Problem Description
Given two integers arrays, A and B, of size N each.
Find the maximum N elements from the sum combinations (Ai + Bj) formed from elements in arrays A and B.

Problem Constraints
1 <= N <= 2 * 105
-1000 <= A[i], B[i] <= 1000

Input Format
The first argument is an integer array A.
The second argument is an integer array B.

Output Format
Return an integer array denoting the N maximum element in descending order.

Example Input
Input 1:
A = [1, 4, 2, 3]
B = [2, 5, 1, 6]

Input 2:
A = [2, 4, 1, 1]
B = [-2, -3, 2, 4]

Example Output
Output 1:
[10, 9, 9, 8]

Output 2:
[8, 6, 6, 5]

Example Explanation
Explanation 1:
4 maximum elements are 10(6+4), 9(6+3), 9(5+4), 8(6+2).

Explanation 2:
4 maximum elements are 8(4+4), 6(4+2), 6(4+2), 5(4+1).
*/
func GetMaxPairCombination(A []int, B []int) []int {
	N := len(A)
	//hash set to ensure that the pair is not duplicated
	hs := make(map[int]bool)

	//sort the input integer arrays
	sort.Ints(A)
	sort.Ints(B)

	ans := make([]int, N)

	//we will use the max heap here
	mh := make(maxStructHeap, 0)
	heap.Init(&mh)

	//push the max sum into the heap with the indexes
	heap.Push(&mh, &info{
		idx1:  N - 1,
		idx2:  N - 1,
		value: A[N-1] + B[N-1],
	})

	//iterate till N-1 times to ensure we don't get arry index out of bounds exception while accessing the array
	for i := 0; i < N-1; i++ {
		elem := heap.Pop(&mh).(*info)

		//store the answer
		ans[i] = elem.value

		//wiser idea would be to use integer key rather than string key
		//to ensure uniqueness, use x*N+y
		key1 := elem.idx1*N + elem.idx2 - 1
		key2 := (elem.idx1-1)*N + elem.idx2
		if !hs[key1] {
			heap.Push(&mh, &info{
				idx1:  elem.idx1,
				idx2:  elem.idx2 - 1,
				value: A[elem.idx1] + B[elem.idx2-1],
			})
			hs[key1] = true
		}
		if !hs[key2] {
			heap.Push(&mh, &info{
				idx1:  elem.idx1 - 1,
				idx2:  elem.idx2,
				value: A[elem.idx1-1] + B[elem.idx2],
			})
			hs[key2] = true
		}
	}

	//save the last element
	ans[N-1] = heap.Pop(&mh).(*info).value
	return ans
}
