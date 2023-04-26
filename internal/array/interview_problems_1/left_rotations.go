package interviewproblems1

/*
Problem Description
Given an array of integers A and multiple values in B, which represents the number of times array A needs to be left rotated.
Find the rotated array for each value and return the result in the from of a matrix where ith row represents the rotated array for the ith value in B.

Problem Constraints
1 <= length of both arrays <= 2000 -10^9 <= A[i] <= 10^9 0 <= B[i] <= 2000

Input Format
The first argument given is the integer array A.
The second argument given is the integer array B.

Output Format
Return the resultant matrix.

Example Input
Input 1:
A = [1, 2, 3, 4, 5]
B = [2, 3]

Input 2:
A = [5, 17, 100, 11]
B = [1]

Example Output
Output 1:
[ [3, 4, 5, 1, 2]

	[4, 5, 1, 2, 3] ]

Output 2:
[ [17, 100, 11, 5] ]

Example Explanation
for input 1 -> B[0] = 2 which requires 2 times left rotations
1: [2, 3, 4, 5, 1]
2: [3, 4, 5, 1, 2]

B[1] = 3 which requires 3 times left rotation
1: [2, 3, 4, 5, 1]
2: [3, 4, 5, 1, 2]
2: [4, 5, 1, 2, 4]
*/
func GetRotatedArray(A []int, B []int) [][]int {
	var result [][]int
	for i := range B {
		localRev := reverse(A, 0, len(A)-1)
		//get the right rotation count from left rotation count
		rightRotationCount := len(A) - (B[i] % len(A))
		//reverse the subarrays
		localRev = reverse(localRev, 0, rightRotationCount-1)
		localRev = reverse(localRev, rightRotationCount, len(A)-1)
		result = append(result, localRev)
	}
	return result
}

func reverse(inp []int, start, end int) []int {
	out := make([]int, len(inp))
	copy(out, inp)
	for i, j := start, end; i <= j; i, j = i+1, j-1 {
		out[i], out[j] = inp[j], inp[i]
	}
	return out
}
