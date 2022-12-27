package sorting1

/*
Problem Description
Given an array of integers A. If i < j and A[i] > A[j], then the pair (i, j) is called an inversion of A. Find the total number of inversions of A modulo (109 + 7).

Problem Constraints
1 <= length of the array <= 100000
1 <= A[i] <= 10^9

Input Format
The only argument given is the integer array A.

Output Format
Return the number of inversions of A modulo (109 + 7).

Example Input
Input 1:
A = [3, 2, 1]

Input 2:
A = [1, 2, 3]

Example Output
Output 1:
3

Output 2:
0

Example Explanation
Explanation 1:
All pairs are inversions.

Explanation 2:
No inversions.
*/

func GetInversionCount(A []int) int {
	return getInversionCount(A, make([]int, len(A)), 0, len(A)-1)
}

func getInversionCount(inp, aux []int, beg, end int) int {
	if beg == end {
		return 0
	}
	var mod int = 1e9 + 7
	mid := (beg + end) / 2
	left := getInversionCount(inp, aux, beg, mid)
	right := getInversionCount(inp, aux, mid+1, end)
	across := mergeSortedArrays(inp, aux, beg, mid, end)
	return (left + right + across) % mod

}

func mergeSortedArrays(inp, sortedArr []int, beg, mid, end int) int {
	idx1, idx2, idx3 := beg, mid+1, beg
	var mod int = 1e9 + 7
	count := 0
	for idx1 <= mid && idx2 <= end {
		if inp[idx1] <= inp[idx2] {
			sortedArr[idx3] = inp[idx1]
			idx1++
		} else {
			count = (count + (mid - idx1 + 1)) % mod
			sortedArr[idx3] = inp[idx2]
			idx2++
		}
		idx3++
	}

	for idx1 <= mid {
		sortedArr[idx3] = inp[idx1]
		idx1++
		idx3++
	}

	for idx2 <= end {
		sortedArr[idx3] = inp[idx2]
		idx2++
		idx3++
	}

	for i := beg; i <= end; i++ {
		inp[i] = sortedArr[i]
	}

	return count
}
