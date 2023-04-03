package dp6

import "sort"

/*
Problem Description
Given a matrix of integers A of size N x 2 describing dimensions of N envelopes, where A[i][0] denotes the height of the ith envelope and A[i][1] denotes the width of the ith envelope.
One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.
Find the maximum number of envelopes you can put one inside other.

Problem Constraints
1 <= N <= 1000
1 <= A[i][0], A[i][1] <= 109

Input Format
The only argument given is the integer matrix A.

Output Format
Return an integer denoting the maximum number of envelopes you can put one inside other.

Example Input
Input 1:
A = [
        [5, 4]
        [6, 4]
        [6, 7]
        [2, 3]
    ]

Input 2:
A = [     '
        [8, 9]
        [8, 18]
    ]

Example Output
Output 1:
3

Output 2:
1

Example Explanation
Explanation 1:
Step 1: put [2, 3] inside [5, 4]
Step 2: put [5, 4] inside [6, 7]
3 envelopes can be put one inside other.

Explanation 2:
No envelopes can be put inside any other so answer is 1.
*/
func GetMaxEnvelopes(A [][]int) int {
	inp := A
	//sort by height first in ASC order and then sort by width in DESC order
	//reason for sorting is atleast we wouldn't have to worry about height
	//when looking at the increasing sequence of width
	sort.Slice(inp, func(i, j int) bool {
		return (inp[i][0] < inp[j][0]) || (inp[i][0] == inp[j][0] && inp[i][1] > inp[j][1])
	})

	return getLIS(inp)

}

func getLIS(inp [][]int) int {
	N := len(inp)
	memo := make([]int, N+1)
	for i := 1; i <= N; i++ {
		maxVal := 0
		//check from o to i-1 th index if there is any smaller values
		for j := 1; j < i; j++ {
			if inp[i-1][0] > inp[j-1][0] && inp[i-1][1] > inp[j-1][1] {
				maxVal = max(maxVal, memo[j])
			}
		}
		memo[i] = maxVal + 1
	}

	ans := 0
	for i := range memo {
		ans = max(ans, memo[i])
	}
	return ans
}
