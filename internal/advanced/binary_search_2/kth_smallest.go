package binarysearch2

import (
	"sort"
)

/*
Problem Description
Given an integer array A of size N.

If we store the sum of each triplet of the array A in a new list, then find the Bth smallest element among the list.

NOTE: A triplet consists of three elements from the array. Let's say if A[i], A[j], A[k] are the elements of the triplet then i < j < k.

Problem Constraints
3 <= N <= 500
1 <= A[i] <= 108
1 <= B <= (N*(N-1)*(N-2))/6

Input Format
The first argument is an integer array A.
The second argument is an integer B.

Output Format
Return an integer denoting the Bth element of the list.

Example Input
Input 1:
A = [2, 4, 3, 2]
B = 3

Input 2:
A = [1, 5, 7, 3, 2]
B = 9

Example Output
Output 1:
9

Output 2:
14

Example Explanation
Explanation 1:
All the triplets of the array A are:
(2, 4, 3) = 9
(2, 4, 2) = 8
(2, 3, 2) = 7
(4, 3, 2) = 9
So the 3rd smallest element is 9.
*/

func GetBthSmallestRank(A []int, B int) int {
	sort.Ints(A)
	var ans, mid int
	arrLength := len(A)
	//range would be from first 3 to last 3 elements
	low := A[0] + A[1] + A[2]
	high := A[arrLength-1] + A[arrLength-2] + A[arrLength-3]
	for low <= high {
		mid = low + (high-low)>>1
		//if the rank is more than B, move left
		if hasMoreRank(A, B, mid) {
			high = mid - 1
			//save ans, move right
		} else {
			ans = mid
			low = mid + 1
		}
	}
	return ans
}

// for the given mid, we will find triplets whose sum is less than mid
// if the triplet count is >=B, it means mid's rank is more than B
func hasMoreRank(inp []int, B, mid int) bool {
	var k, val int
	var count int
	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
			val = mid - (inp[i] + inp[j])
			if val >= inp[j] {
				k = doBS(inp, val, j+1)
				if k != -1 {
					count += k - j
				}
			} else {
				break
			}
		}
	}

	return (count) >= B
}

// returns the element just less than val
func doBS(inp []int, val, startIdx int) int {
	k := -1
	low := startIdx
	high := len(inp) - 1
	var mid int
	for low <= high {
		mid = low + (high-low)>>1
		if inp[mid] < val {
			k = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return k
}
