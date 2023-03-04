package heaps1

/*
Problem Description
Given a sorted matrix of integers A of size N x M and an integer B.
Each of the rows and columns of matrix A is sorted in ascending order, find the Bth smallest element in the matrix.
NOTE: Return The Bth smallest element in the sorted order, not the Bth distinct element.

Problem Constraints
1 <= N, M <= 500
1 <= A[i] <= 109
1 <= B <= N * M

Input Format
The first argument given is the integer matrix A.
The second argument given is an integer B.

Output Format
Return the B-th smallest element in the matrix.

Example Input
Input 1:
A = [ [9, 11, 15],
      [10, 15, 17] ]
B = 6

Input 2:
A = [  [5, 9, 11],
       [9, 11, 13],
       [10, 12, 15],
       [13, 14, 16],
       [16, 20, 21] ]
B = 12

Example Output
Output 1:
17

Output 2:
16

Example Explanation
Explanation 1:
6th smallest element in the sorted matrix is 17.

Explanation 2:
12th smallest element in the sorted matrix is 16.
*/
func GetKthSmallestElem(A [][]int, B int) int {
	var heap []int
	for i := range A {
		for j := range A[i] {
			//populate till heap size is full
			if len(heap) < B {
				heap = append(heap, A[i][j])
				//once full, create a max heap
				if len(heap) == B {
					createMaxHeap(heap)
				}
				//if the cur element is smaller than the max, remove the max, insert the cur
			} else if heap[0] > A[i][j] {
				heap[0] = A[i][j]
				maxHeapify(0, heap)
			}
		}
	}
	return heap[0]
}
