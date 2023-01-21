package hashing1

/*
Problem Description
Given an unsorted integer array A of size N.
Find the length of the longest set of consecutive elements from array A.

Problem Constraints
1 <= N <= 106
-106 <= A[i] <= 106

Input Format
First argument is an integer array A of size N.

Output Format
Return an integer denoting the length of the longest set of consecutive elements from the array A.

Example Input
Input 1:
A = [100, 4, 200, 1, 3, 2]

Input 2:
A = [2, 1]

Example Output
Output 1:
4

Output 2:
2

Example Explanation
Explanation 1:
The set of consecutive elements will be [1, 2, 3, 4].

Explanation 2:
The set of consecutive elements will be [1, 2].
*/
func LongestConsecutive(A []int) int {
	maxCount, count := 0, 0
	hashSet := make(map[int]bool)
	for i := range A {
		hashSet[A[i]] = true
	}

	for k := range hashSet {
		//if k is the start element of the sequence
		if !hashSet[k-1] {
			count = 0
			start := k
			for hashSet[start] {
				count++
				start++
			}
			if count > maxCount {
				maxCount = count
			}
		}
	}
	return maxCount
}
