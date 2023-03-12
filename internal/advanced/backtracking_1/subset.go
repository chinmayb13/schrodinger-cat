package backtracking1

import (
	"math"
	"sort"
)

/*
Problem Description
Given a set of distinct integers A, return all possible subsets.

NOTE:
Elements in a subset must be in non-descending order.
The solution set must not contain duplicate subsets.
Also, the subsets should be sorted in ascending ( lexicographic ) order.
The list is not necessarily sorted.

Problem Constraints
1 <= |A| <= 16
INTMIN <= A[i] <= INTMAX

Input Format
First and only argument of input contains a single integer array A.

Output Format
Return a vector of vectors denoting the answer.

Example Input
Input 1:
A = [1]

Input 2:
A = [1, 2, 3]

Example Output
Output 1:
[
    []
    [1]
]

Output 2:
[
 []
 [1]
 [1, 2]
 [1, 2, 3]
 [1, 3]
 [2]
 [2, 3]
 [3]
]

Example Explanation
Explanation 1:
You can see that these are all possible subsets.

Explanation 2:
You can see that these are all possible subsets.
*/

func SubSets(A []int) [][]int {
	sort.Ints(A)
	var cur []int
	var ans [][]int
	//generate all the subsets
	generateSubSet(A, &ans, cur, 0)

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

func generateSubSet(inp []int, ans *[][]int, cur []int, idx int) {
	if idx == len(inp) {
		*ans = append(*ans, copyArr(cur))
		return
	}
	generateSubSet(inp, ans, cur, idx+1)
	cur = append(cur, inp[idx])
	generateSubSet(inp, ans, cur, idx+1)
	cur = cur[:len(cur)-1]

}

func min(args ...int) int {
	minVal := math.MaxInt
	for i := range args {
		if args[i] < minVal {
			minVal = args[i]
		}
	}
	return minVal
}
