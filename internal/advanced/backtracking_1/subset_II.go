package backtracking1

import "sort"

/*
Problem Description
Given a collection of integers denoted by array A of size N that might contain duplicates, return all possible subsets.

NOTE:
Elements in a subset must be in non-descending order.
The solution set must not contain duplicate subsets.
The subsets must be sorted lexicographically.

Problem Constraints
0 <= N <= 16

Input Format
Only argument is an integer array A of size N.

Output Format
Return a 2-D vector denoting all the possible subsets.

Example Input
Input 1:
A = [1, 2, 2]

Input 2:
A = [1, 1]

Example Output
Output 1:
[

	[],
	[1],
	[1, 2],
	[1, 2, 2],
	[2],
	[2, 2]

]

Output 2:
[

	[],
	[1],
	[1, 1]

]

Example Explanation
Explanation 1:
All the subsets of the array [1, 2, 2] in lexicographically sorted order.
*/
func SubsetsWithDup(A []int) [][]int {
	sort.Ints(A)
	var cur []int
	var ans [][]int
	//generate all the subsets
	generateSubSetWithDup(A, &ans, cur, 0)

	//sort the subsets
	sort.Slice(ans, func(i, j int) bool {
		for k := 0; k < min(len(ans[i]), len(ans[j])); k++ {
			if ans[i][k] != ans[j][k] {
				return ans[i][k] < ans[j][k]

			}
		}
		return len(ans[i]) < len(ans[j])
	})
	return ans
}

func generateSubSetWithDup(inp []int, ans *[][]int, cur []int, idx int) {
	if idx == len(inp) {
		*ans = append(*ans, copyArr(cur))
		return
	}

	//include the element
	cur = append(cur, inp[idx])
	generateSubSetWithDup(inp, ans, cur, idx+1)

	//exclude the element
	cur = cur[:len(cur)-1]

	//since the next idx could also contain the same element
	//increase the idx such that idx+1 contains different element
	for (idx+1) < len(inp) && inp[idx] == inp[idx+1] {
		idx++
	}
	generateSubSetWithDup(inp, ans, cur, idx+1)
}
