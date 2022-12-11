package bitmanipulate2

/*
Problem Description
Given an array of positive integers A, two integers appear only once, and all the other integers appear twice.
Find the two integers that appear only once.

Note: Return the two numbers in ascending order.

Problem Constraints
2 <= |A| <= 100000
1 <= A[i] <= 109

Input Format
The first argument is an array of integers of size N.

Output Format
Return an array of two integers that appear only once.

Example Input
Input 1:
A = [1, 2, 3, 1, 2, 4]

Input 2:
A = [1, 2]

Example Output
Output 1:
[3, 4]

Output 2:
[1, 2]

Example Explanation
Explanation 1:
3 and 4 appear only once.

Explanation 2:
1 and 2 appear only once.
*/
func GetTwiceAppearing(A []int) []int {
	xor := 0
	for i := range A {
		xor ^= A[i]
	}

	position := -1
	for j := 0; j < 32; j++ {
		if (xor & (1 << j)) > 0 {
			position = j
			break
		}
	}
	bucketSet := 0
	bucketUnset := 0
	for i := range A {
		if (A[i] & (1 << position)) > 0 {
			bucketSet ^= A[i]
		} else {
			bucketUnset ^= A[i]
		}
	}

	if bucketSet < bucketUnset {
		return []int{bucketSet, bucketUnset}
	}
	return []int{bucketUnset, bucketSet}
}
