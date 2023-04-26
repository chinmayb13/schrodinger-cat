package interviewproblems2

/*
Problem Description
Given an array of size N, find the majority element. The majority element is the element that appears more than floor(n/2) times.
You may assume that the array is non-empty and the majority element always exists in the array.

Problem Constraints
1 <= N <= 5*105
1 <= num[i] <= 109

Input Format
Only argument is an integer array.

Output Format
Return an integer.

Example Input
[2, 1, 2]

Example Output
2

Example Explanation
2 occurs 2 times which is greater than 3/2.
*/
func GetMajorityElement(A []int) int {
	cme := 0
	frequency := 0
	for i := range A {
		if frequency == 0 {
			cme = A[i]
			frequency = 1
		} else if cme == A[i] {
			frequency++
		} else {
			frequency--
		}
	}
	return cme
}
