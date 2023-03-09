package greedyalgo

/*
Problem Description
N children are standing in a line. Each child is assigned a rating value.
You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
What is the minimum number of candies you must give?

Problem Constraints
1 <= N <= 105
-109 <= A[i] <= 109

Input Format
The first and only argument is an integer array A representing the rating of children.

Output Format
Return an integer representing the minimum candies to be given.

Example Input
Input 1:
A = [1, 2]

Input 2:
A = [1, 5, 2, 1]

Example Output
Output 1:
3

Output 2:
7

Example Explanation
Explanation 1:
The candidate with 1 rating gets 1 candy and candidate with rating 2 cannot get 1 candy as 1 is its neighbor.
So rating 2 candidate gets 2 candies. In total, 2 + 1 = 3 candies need to be given out.

Explanation 2:
Candies given = [1, 3, 2, 1]
*/
//idea : calculate min from left and right side and take the max
func GetMinCandies(A []int) int {
	ans := 0
	N := len(A)
	leftCalculation := leftView(A)
	rightCalculation := rightView(A)
	for i := 0; i < N; i++ {
		ans += max(leftCalculation[i], rightCalculation[i])
	}
	return ans
}

func max(args ...int) int {
	ans := 0
	for i := range args {
		if args[i] > ans {
			ans = args[i]
		}
	}
	return ans
}

func leftView(inp []int) []int {
	N := len(inp)
	ans := make([]int, N)
	ans[0] = 1
	for i := 1; i < N; i++ {
		if inp[i] > inp[i-1] {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = 1
		}
	}
	return ans
}

func rightView(inp []int) []int {
	N := len(inp)
	ans := make([]int, N)
	ans[N-1] = 1
	for i := N - 2; i >= 0; i-- {
		if inp[i] > inp[i+1] {
			ans[i] = ans[i+1] + 1
		} else {
			ans[i] = 1
		}
	}
	return ans
}
