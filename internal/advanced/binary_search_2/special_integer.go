package binarysearch2

/*
Problem Description
Given an array of integers A and an integer B, find and return the maximum value K such that there is no subarray in A of size K with the sum of elements greater than B.

Problem Constraints
1 <= |A| <= 100000
1 <= A[i] <= 10^9
1 <= B <= 10^9

Input Format
The first argument given is the integer array A.
The second argument given is integer B.

Output Format
Return the maximum value of K (sub array length).

Example Input
Input 1:
A = [1, 2, 3, 4, 5]
B = 10

Input 2:
A = [5, 17, 100, 11]
B = 130

Example Output
Output 1:
2

Output 2:
3

Example Explanation
Explanation 1:
Constraints are satisfied for maximal value of 2.

Explanation 2:
Constraints are satisfied for maximal value of 3.
*/
func GetSpecialInteger(A []int, B int) int {
	ans := 0
	low := 1
	high := len(A)
	for low <= high {
		mid := low + (high-low)>>1
		if doesSubArrayExist(A, mid, int64(B)) {
			high = mid - 1
		} else {
			ans = mid
			low = mid + 1
		}
	}
	return ans
}

func doesSubArrayExist(inp []int, size int, val int64) bool {
	var sum int64
	for i := 0; i < size; i++ {
		sum += int64(inp[i])
	}
	if sum > val {
		return true
	}
	l := 1
	r := size
	for r < len(inp) {
		sum = sum + int64(inp[r]-inp[l-1])
		if sum > val {
			return true
		}
		l++
		r++
	}
	return false
}
