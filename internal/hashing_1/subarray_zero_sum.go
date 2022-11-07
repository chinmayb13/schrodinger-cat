package hashing1

/*
Problem Description
Given an array of integers A, find and return whether the given array contains a non-empty subarray with a sum equal to 0.
If the given array contains a sub-array with sum zero return 1, else return 0.

Problem Constraints
1 <= |A| <= 100000
-10^9 <= A[i] <= 10^9

Input Format
The only argument given is the integer array A.

Output Format
Return whether the given array contains a subarray with a sum equal to 0.

Example Input
Input 1:
A = [1, 2, 3, 4, 5]

Input 2:
A = [-1, 1]

Example Output
Output 1:
0

Output 2:
1

Example Explanation
Explanation 1:
No subarray has sum 0.

Explanation 2:
The array has sum 0.
*/
//approach: take cumulative sum at each index and check if the same sum has been calculated before. if it is. then return 1
//idea: sum of subarray
//=> sum[L,R]=0
//=> sum[0,R]-sum[0,L-1]=0
//=> sum[0,R] = sum[0,L-1]
//same sum is appearing twice
func HasZeroSumSubarray(A []int) int {
	sumMap := map[int]bool{0: true}
	sum := 0
	for i := range A {
		sum += A[i]
		if sumMap[sum] {
			return 1
		}
		sumMap[sum] = true
	}
	return 0
}
