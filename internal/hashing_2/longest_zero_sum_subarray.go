package hashing2

/*
Problem Description
Given an array A of N integers.
Find the length of the longest subarray in the array which sums to zero.

Problem Constraints
1 <= N <= 105
-109 <= A[i] <= 109

Input Format
Single argument which is an integer array A.

Output Format
Return an integer.

Example Input
Input 1:
A = [1, -2, 1, 2]

Input 2:
A = [3, 2, -1]

Example Output
Output 1:
3

Output 2:
0

Example Explanation
Explanation 1:
[1, -2, 1] is the largest subarray which sums up to 0.

Explanation 2:
No subarray sums up to 0.
*/
func GetLongestSubArrayLength(A []int) int {
	hashMap := make(map[int]int)
	length := 0
	sum := 0
	for i := range A {
		sum += A[i]
		//if subarray length is zero
		if sum == 0 && (i+1) > length {
			length = i + 1
		}

		idx, ok := hashMap[sum]
		//if sum is repeating then the sandwich subarray has zero sum
		if ok && (i-idx) > length {
			length = i - idx
		} else if !ok {
			hashMap[sum] = i
		}

	}
	return length
}
