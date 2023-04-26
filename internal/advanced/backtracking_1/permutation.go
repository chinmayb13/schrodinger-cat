package backtracking1

/*
Problem Description
Given an integer array A of size N denoting collection of numbers , return all possible permutations.

NOTE:
No two entries in the permutation sequence should be the same.
For the purpose of this problem, assume that all the numbers in the collection are unique.
Return the answer in any order
WARNING: DO NOT USE LIBRARY FUNCTION FOR GENERATING PERMUTATIONS.
Example : next_permutations in C++ / itertools.permutations in python.
If you do, we will disqualify your submission retroactively and give you penalty points.

Problem Constraints
1 <= N <= 9

Input Format
Only argument is an integer array A of size N.

Output Format
Return a 2-D array denoting all possible permutation of the array.

Example Input
A = [1, 2, 3]

Example Output
[ [1, 2, 3]

	[1, 3, 2]
	[2, 1, 3]
	[2, 3, 1]
	[3, 1, 2]
	[3, 2, 1] ]

Example Explanation
All the possible permutation of array [1, 2, 3].
*/
func GeneratePermute(A []int) [][]int {
	hs := make(map[int]bool)
	cur := make([]int, len(A))
	var ans [][]int
	generatePerm(A, 0, hs, cur, &ans)
	return ans
}

func generatePerm(inp []int, idx int, hs map[int]bool, cur []int, ans *[][]int) {
	if idx == len(cur) {
		*ans = append(*ans, copyArr(cur))
		return
	}

	for i := range inp {
		if !hs[inp[i]] {
			cur[idx] = inp[i]
			hs[inp[i]] = true
			generatePerm(inp, idx+1, hs, cur, ans)
			delete(hs, inp[i])
		}
	}
}
