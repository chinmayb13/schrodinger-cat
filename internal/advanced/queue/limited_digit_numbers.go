package queue

/*
Problem Description
Given an integer, A. Find and Return first positive A integers in ascending order containing only digits 1, 2, and 3.
NOTE: All the A integers will fit in 32-bit integers.

Problem Constraints
1 <= A <= 29500

Input Format
The only argument given is integer A.

Output Format
Return an integer array denoting the first positive A integers in ascending order containing only digits 1, 2 and 3.

Example Input
Input 1:
A = 3

Input 2:
A = 7

Example Output
Output 1:
[1, 2, 3]

Output 2:
[1, 2, 3, 11, 12, 13, 21]

Example Explanation
Explanation 1:
Output denotes the first 3 integers that contains only digits 1, 2 and 3.

Explanation 2:
Output denotes the first 3 integers that contains only digits 1, 2 and 3.
*/
func GetNumbersWithLimitedDigits(A int) []int {
	ans := make([]int, A+2)
	ans[0] = 1
	ans[1] = 2
	ans[2] = 3
	front, back := 0, 2
	for back < A-1 {
		target := ans[front] * 10
		//for ans[front]=1, below loop generates 11,12,13
		for i := 0; i < 3; i++ {
			back++
			target++
			ans[back] = target
		}
		front++
	}
	return ans[:A]
}
