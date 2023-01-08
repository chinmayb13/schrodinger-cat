package sorting3

/*
Problem Description
Given an array A. Sort this array using Count Sort Algorithm and return the sorted array.

Problem Constraints
1 <= |A| <= 105
1 <= A[i] <= 105

Input Format
The first argument is an integer array A.

Output Format
Return an integer array that is the sorted array A.

Example Input
Input 1:
A = [1, 3, 1]

Input 2:
A = [4, 2, 1, 3]

Example Output
Output 1:
[1, 1, 3]

Output 2:
[1, 2, 3, 4]

Example Explanation
For Input 1:
The array in sorted order is [1, 1, 3].

For Input 2:
The array in sorted order is [1, 2, 3, 4].
*/
func CountSort(A []int) []int {
	hashMap := make(map[int]int)
	for i := range A {
		hashMap[A[i]]++
	}
	k := 0
	for i := 1; i <= 1e5; i++ {
		freq := hashMap[i]
		for j := 1; j <= freq; j++ {
			A[k] = i
			k++
		}
	}
	return A
}
