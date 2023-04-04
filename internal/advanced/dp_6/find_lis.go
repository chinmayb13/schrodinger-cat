package dp6

/*
Problem Description
You are given an array A. You need to find the length of the Longest Increasing Subsequence in the array.
In other words, you need to find a subsequence of array A in which the elements are in sorted order, (strictly increasing) and as long as possible.

Problem Constraints
1 ≤ length(A), A[i] ≤ 105

Input Format
The first and only argument of the input is the array A.

Output Format
Output a single integer, the length of the longest increasing subsequence in array A.

Example Input
Input 1:
A: [2, 1, 4, 3]

Input 2:
A: [5, 6, 3, 7, 9]

Example Output
Output 1:
2

Output 2:
4

Example Explanation
Explanation 1:
[2, 4] and [1, 3] are the longest increasing sequences of size 2.

Explanation 2:
The longest increasing subsequence we can get is [5, 6, 7, 9] of size 4.
*/
func FindLIS(A []int) int {
	inp := A
	size := 1
	tails := make([]int, len(inp))
	for i := range inp {
		//if element is greater than the greatest element in the tails
		//add at the end, increase the size
		if inp[i] > tails[size-1] {
			tails[size] = inp[i]
			size++
			continue
		}
		//else save the element at the position i
		//where A[i]>=element
		low := 0
		high := size
		var idx int
		for low <= high {
			mid := low + (high-low)>>1
			if tails[mid] >= inp[i] {
				idx = mid
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		tails[idx] = inp[i]
		if idx == size {
			size++
		}
	}
	return size - 1
}
