package binarysearch1

/*
Problem Description
Given a sorted array of integers A(0 based index) of size N, find the starting and the ending position of a given integer B in array A.
Your algorithm's runtime complexity must be in the order of O(log n).
Return an array of size 2, such that the first element = starting position of B in A and the second element = ending position of B in A, if B is not found in A return [-1, -1].

Problem Constraints
1 <= N <= 106
1 <= A[i], B <= 109

Input Format
The first argument given is the integer array A.
The second argument given is the integer B.

Output Format
Return an array of size 2, such that the first element = starting position of B in A and the second element = the ending position of B in A if B is not found in A return [-1, -1].

Example Input
Input 1:
A = [5, 7, 7, 8, 8, 10]
B = 8

Input 2:
A = [5, 17, 100, 111]
B = 3

Example Output
Output 1:
[3, 4]

Output 2:
[-1, -1]

Example Explanation
Explanation 1:
The first occurence of 8 in A is at index 3.
The second occurence of 8 in A is at index 4.
ans = [3, 4]

Explanation 2:
There is no occurence of 3 in the array.
*/
func SearchRange(A []int, B int) []int {
	return []int{getStartIdx(A, B), getEndIdx(A, B)}
}

func getStartIdx(inp []int, val int) int {
	low := 0
	high := len(inp) - 1
	var ans int
	for low <= high {
		idx := low + (high-low)/2
		//this condition will always return the first occurence
		if inp[idx] >= val {
			ans = idx
			high = idx - 1
		} else {
			low = idx + 1
		}
	}
	if inp[ans] == val {
		return ans
	}
	return -1

}

func getEndIdx(inp []int, val int) int {
	low := 0
	high := len(inp) - 1
	var ans int
	for low <= high {
		idx := low + (high-low)/2
		//this condition will always return the last occurence
		if inp[idx] <= val {
			ans = idx
			low = idx + 1
		} else {
			high = idx - 1
		}
	}
	if inp[ans] == val {
		return ans
	}
	return -1

}
