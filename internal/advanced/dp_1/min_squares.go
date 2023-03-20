package dp1

import "math"

/*
Problem Description
Given an integer A. Return minimum count of numbers, sum of whose squares is equal to A.

Problem Constraints
1 <= A <= 105

Input Format
First and only argument is an integer A.

Output Format
Return an integer denoting the minimum count.

Example Input
Input 1:
A = 6

Input 2:
A = 5

Example Output
Output 1:
3

Output 2:
2

Example Explanation
Explanation 1:
Possible combinations are : (12 + 12 + 12 + 12 + 12 + 12) and (12 + 12 + 22).
Minimum count of numbers, sum of whose squares is 6 is 3.

Explanation 2:
We can represent 5 using only 2 numbers i.e. 12 + 22 = 5
*/
func CountMinSquares(A int) int {
	memo := make([]int, A+1)
	for i := range memo {
		memo[i] = -1
	}
	return minCount(A, memo)
}

func minCount(N int, memo []int) int {
	//base condition
	if N == 0 {
		return 0
	}

	//check
	if memo[N] != -1 {
		return memo[N]
	}

	//calculate
	min := math.MaxInt
	sqrt := int(math.Sqrt(float64(N)))
	for i := 1; i <= sqrt; i++ {
		val := minCount(N-i*i, memo)
		if val < min {
			min = val
		}
	}
	min++

	//store
	memo[N] = min

	//return
	return min
}
