package heaps2

import "container/heap"

/*
Problem Description
Given an array A of N numbers, you have to perform B operations. In each operation, you have to pick any one of the N elements and add the original value(value stored at the index before we did any operations) to its current value. You can choose any of the N elements in each operation.

Perform B operations in such a way that the largest element of the modified array(after B operations) is minimized.
Find the minimum possible largest element after B operations.

Problem Constraints
1 <= N <= 106
0 <= B <= 105
-105 <= A[i] <= 105

Input Format
The first argument is an integer array A.
The second argument is an integer B.

Output Format
Return an integer denoting the minimum possible largest element after B operations.

Example Input
Input 1:
A = [1, 2, 3, 4]
B = 3

Input 2:
A = [5, 1, 4, 2]
B = 5

Example Output
Output 1:
4

Output 2:
5

Example Explanation
Explanation 1:
Apply operation on element at index 0, the array would change to [2, 2, 3, 4]
Apply operation on element at index 0, the array would change to [3, 2, 3, 4]
Apply operation on element at index 0, the array would change to [4, 2, 3, 4]
Minimum possible largest element after 3 operations is 4.

Explanation 2:
Apply operation on element at index 1, the array would change to [5, 2, 4, 2]
Apply operation on element at index 1, the array would change to [5, 3, 4, 2]
Apply operation on element at index 1, the array would change to [5, 4, 4, 2]
Apply operation on element at index 1, the array would change to [5, 5, 4, 2]
Apply operation on element at index 3, the array would change to [5, 5, 4, 4]
Minimum possible largest element after 5 operations is 5.
*/
func GetMinLargestElem(A []int, B int) int {
	mh := make(minStructHeap, 0)
	var max int = -1e5 - 1
	heap.Init(&mh)
	for i := range A {
		if A[i] > max {
			max = A[i]
		}
		//keeping pushing the elements into min Heap
		//the Less() func automatically arranges the elements in a fashion
		//the elem with minimum cur value + original value would be at the top
		heap.Push(&mh, &minInfo{
			value: A[i],
			inc:   A[i],
		})
	}

	//do B operations
	for B > 0 {
		elem := heap.Pop(&mh).(*minInfo)
		elem.value += elem.inc
		if elem.value > max {
			max = elem.value
		}
		heap.Push(&mh, elem)
		B--
	}
	return max

}

type minInfo struct {
	value int
	inc   int
}

type minStructHeap []*minInfo

func (h minStructHeap) Len() int { return len(h) }
func (h minStructHeap) Less(i, j int) bool {
	//sort by minimum value taken after incrementing the current value with its original value
	return h[i].value+h[i].inc < h[j].value+h[j].inc
}
func (h minStructHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minStructHeap) Push(x interface{}) {
	*h = append(*h, x.(*minInfo))
}

func (h *minStructHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}
