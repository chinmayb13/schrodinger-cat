package hashing1

/*
Problem Description
You are given an array A of N integers. Return the count of elements with frequncy 1 in the given array.

Problem Constraints
1 <= N <= 105
1 <= A[i] <= 109

Input Format
First argument A is an array of integers.

Output Format
Return an integer.

Example Input
Input 1:
A = [3, 4, 3, 6, 6]

Input 2:
A = [3, 3, 3, 9, 0, 1, 0]

Example Output
Output 1:
1

Output 2:
2


Example Explanation
For Input 1:
Only the element 4 has a frequency 1.

For Input 2:
The elements 9 and 1 has a frequncy 1.
*/
func GetSingleFreqCount(A []int) int {
	presenceMap := make(map[int]int)
	count := 0
	for i := range A {
		f, ok := presenceMap[A[i]]
		if ok && f > 0 {
			count--
			presenceMap[A[i]] = 0
		} else if !ok {
			presenceMap[A[i]] = 1
			count++
		}
	}
	return count
}
