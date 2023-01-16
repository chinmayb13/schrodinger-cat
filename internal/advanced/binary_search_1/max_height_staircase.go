package binarysearch1

/*
Problem Description
Given an integer A representing the number of square blocks. The height of each square block is 1. The task is to create a staircase of max-height using these blocks.
The first stair would require only one block, and the second stair would require two blocks, and so on.
Find and return the maximum height of the staircase.

Problem Constraints
0 <= A <= 109

Input Format
The only argument given is integer A.

Output Format
Return the maximum height of the staircase using these blocks.

Example Input
Input 1:
A = 10

Input 2:
A = 20

Example Output
Output 1:
4

Output 2:
5

Example Explanation
Explanation 1:
The stairs formed will have height 1, 2, 3, 4.

Explanation 2:
The stairs formed will have height 1, 2, 3, 4, 5.
*/
func GetMaxHeight(A int) int {
	var ans int64
	var low int64 = 0
	var high int64 = int64(A)
	for low <= high {
		mid := low + (high-low)>>1
		if ((mid * (mid + 1)) >> 1) <= int64(A) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return int(ans)
}

func GetMaxHeightAlt(A int) int {
	height := 0
	for i := 0; A > 0; i++ {
		A -= (i + 1)
		height++
	}
	if A < 0 {
		return height - 1
	}
	return height
}
