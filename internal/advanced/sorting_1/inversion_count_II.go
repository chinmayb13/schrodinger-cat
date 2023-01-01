package sorting1

/*
Problem Description
Given an array of integers A, we call (i, j) an important reverse pair if i < j and A[i] > 2*A[j].
Return the number of important reverse pairs in the given array A.

Problem Constraints
1 <= length of the array <= 105
-2 * 109 <= A[i] <= 2 * 109

Input Format
The only argument given is the integer array A.

Output Format
Return the number of important reverse pairs in the given array A.

Example Input
Input 1:
A = [1, 3, 2, 3, 1]

Input 2:
A = [4, 1, 2]

Example Output
Output 1:
2

Output 2:
1

Example Explanation
Explanation 1:
There are two pairs which are important reverse i.e (3, 1) and (3, 1).

Explanation 2:
There is only one pair i.e (4, 1).
*/

func GetInversionCountII(A []int) int {
	aux := make([]int, len(A))
	return getInversionCountII(A, aux, 0, len(A)-1)
}

func getInversionCountII(inp, aux []int, beg, end int) int {
	if beg == end {
		return 0
	}
	mid := (beg + end) / 2
	left := getInversionCountII(inp, aux, beg, mid)
	right := getInversionCountII(inp, aux, mid+1, end)
	across := mergeSortedArraysII(inp, aux, beg, mid, end)
	return left + right + across
}

func mergeSortedArraysII(inp, aux []int, beg, mid, end int) int {
	ans := 0
	idx1, idx2 := beg, mid+1
	//since array can contain negative element, therefore comparing them beforehand with the formula would be correct
	//as -530 < -512 but -530 > 2*(-512)
	for idx1 <= mid && idx2 <= end {
		if inp[idx1] > (2 * inp[idx2]) {
			ans += mid - idx1 + 1
			idx2++
		} else {
			idx1++
		}
	}

	idx1, idx2, idx3 := beg, mid+1, beg
	for idx1 <= mid && idx2 <= end {
		if inp[idx1] <= inp[idx2] {
			aux[idx3] = inp[idx1]
			idx1++
		} else {
			aux[idx3] = inp[idx2]
			idx2++
		}
		idx3++
	}

	for idx1 <= mid {
		aux[idx3] = inp[idx1]
		idx1++
		idx3++
	}

	for idx2 <= end {
		aux[idx3] = inp[idx2]
		idx2++
		idx3++
	}

	for i := beg; i <= end; i++ {
		inp[i] = aux[i]
	}

	return ans
}
