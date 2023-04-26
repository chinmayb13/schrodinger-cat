package dp4

import "math"

/*
Given two integer arrays A and B of size N each which represent values and weights associated with N items respectively.
Also given an integer C which represents knapsack capacity.
Find out the maximum value subset of A such that sum of the weights of this subset is smaller than or equal to C.
NOTE: You cannot break an item, either pick the complete item, or don’t pick it (0-1 property).

Problem Constraints
1 <= N <= 500

1 <= C, B[i] <= 106

1 <= A[i] <= 50

Input Format
First argument is an integer array A of size N denoting the values on N items.

Second argument is an integer array B of size N denoting the weights on N items.

Third argument is an integer C denoting the knapsack capacity.

Output Format
Return a single integer denoting the maximum value subset of A such that sum of the weights of this subset is smaller than or equal to C.

Example Input
Input 1:

	A = [6, 10, 12]
	B = [10, 20, 30]
	C = 50

Input 2:

	A = [1, 3, 2, 4]
	B = [12, 13, 15, 19]
	C = 10

Example Output
Output 1:

	22

Output 2:

	0

Example Explanation
Explanation 1:

	Taking items with weight 20 and 30 will give us the maximum value i.e 10 + 12 = 22

Explanation 2:

	Knapsack capacity is 10 but each item has weight greater than 10 so no items can be considered in the knapsack therefore answer is 0.
*/
func GetMax01KnapSackII(A []int, B []int, C int) int {
	N := len(A)
	rows := N + 1
	cols := 50*N + 1
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
		//keep the 0th col of every row as 0
		//because to store 0 profit, min capacity needed is 0
		for j := 1; j < len(memo[i]); j++ {
			memo[i][j] = math.MaxInt
		}
	}

	for i := 1; i <= rows-1; i++ {
		for j := 0; j <= cols-1; j++ {
			//if we don't pick the item
			if memo[i-1][j] != math.MaxInt {
				memo[i][j] = min(memo[i][j], memo[i-1][j])
			}

			//if we pick the item
			if j >= A[i-1] && memo[i-1][j-A[i-1]] != math.MaxInt {
				memo[i][j] = min(memo[i][j], B[i-1]+memo[i-1][j-A[i-1]])
			}
		}
	}

	//return greatest capapcity smaller than equal to C
	for i := cols - 1; i >= 0; i-- {
		if memo[N][i] <= C {
			return i
		}
	}

	return 0
}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}
