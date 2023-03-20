package dp1

/*
Problem Description

Given a 2 x N grid of integer, A, choose numbers such that the sum of the numbers is maximum and no two chosen numbers are adjacent horizontally, vertically or diagonally, and return it.
Note: You can choose more than 2 numbers.

Problem Constraints
1 <= N <= 20000
1 <= A[i] <= 2000

Input Format
The first and the only argument of input contains a 2d matrix, A.

Output Format
Return an integer, representing the maximum possible sum.

Example Input
Input 1:
A = [
       [1]
       [2]
    ]

Input 2:
A = [
       [1, 2, 3, 4]
       [2, 3, 4, 5]
    ]

Example Output
Output 1:
2

Output 2:
8

Example Explanation
Explanation 1:
We will choose 2.

Explanation 2:
We will choose 3 and 5.
*/
func MaxNonAdjacent(A [][]int) int {
	memo := make([]int, len(A[0]))
	for i := range memo {
		memo[i] = -1
	}
	//do it for 0th index
	return calcMax(0, A, memo)
}

func calcMax(idx int, inp [][]int, memo []int) int {
	//if index is out of bounds
	//return 0
	if idx >= len(inp[0]) {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	ans := 0
	//get the max of both rows for current column index
	maxAtCurIdx := max(inp[0][idx], inp[1][idx])

	//if we are including the current element, we need to look for idx+2
	//else we need to look for idx+1 otherwise the element would be adjacent
	//take the max of either of them
	ans += max(maxAtCurIdx+calcMax(idx+2, inp, memo), calcMax(idx+1, inp, memo))
	memo[idx] = ans
	return ans
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}
