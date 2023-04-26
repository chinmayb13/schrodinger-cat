package dp4

/*
Problem Description
Given a strictly increasing array A of positive integers forming a sequence.

# A sequence X1, X2, X3, ..., XN is fibonacci like if

N > =3
Xi + Xi+1 = Xi+2 for all i+2 <= N
Find and return the length of the longest Fibonacci-like subsequence of A.
If one does not exist, return 0.

NOTE: A subsequence is derived from another sequence A by deleting any number of elements (including none) from A, without changing the order of the remaining elements.

Problem Constraints
3 <= length of the array <= 1000
1 <= A[i] <= 109

Input Format
The only argument given is the integer array A.

Output Format
Return the length of the longest Fibonacci-like subsequence of A.
If one does not exist, return 0.

Example Input
Input 1:
A = [1, 2, 3, 4, 5, 6, 7, 8]

Input 2:
A = [1, 3, 7, 11, 12, 14, 18]

Example Output
Output 1:
5

Output 2:
3

Example Explanation
Explanation 1:
The longest subsequence that is fibonacci-like: [1, 2, 3, 5, 8].

Explanation 2:
The longest subsequence that is fibonacci-like: [1, 11, 12], [3, 11, 14] or [7, 11, 18].
The length will be 3.
*/
func GetLongestFibSequence(A []int) int {
	hashMap := make(map[int]int)
	inp := A
	N := len(inp)
	for i := range inp {
		hashMap[inp[i]] = i
	}

	memo := make([][]int, N)
	for i := range memo {
		memo[i] = make([]int, N)
		//every pair will atleast contain 2 items
		for j := range memo[i] {
			memo[i][j] = 2
		}
	}
	max := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			//if the sum exists, increase the count in [j,idx]
			if idx, ok := hashMap[inp[i]+inp[j]]; ok {
				memo[j][idx] = memo[i][j] + 1
				if memo[j][idx] > 2 && memo[j][idx] > max {
					max = memo[j][idx]
				}
			}
		}
	}
	return max
}
