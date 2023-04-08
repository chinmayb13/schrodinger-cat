package dp7

import "math"

/*
Problem Description
Say you have an array, A, for which the ith element is the price of a given stock on day i.
If you were only permitted to complete at most one transaction (ie, buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Return the maximum possible profit.

Problem Constraints
0 <= len(A) <= 7e5
1 <= A[i] <= 1e7

Input Format
The first and the only argument is an array of integers, A.

Output Format
Return an integer, representing the maximum possible profit.

Example Input
Input 1:
A = [1, 2]

Input 2:
A = [1, 4, 5, 2, 4]

Example Output
Output 1:
1

Output 2:
4

Example Explanation
Explanation 1:
Buy the stock on day 0, and sell it on day 1.

Explanation 2:
Buy the stock on day 0, and sell it on day 2.
*/
func MaxProfit(A []int) int {
	inp := A
	buyPrice := math.MaxInt
	ans := 0
	for i := range inp {
		//check if stock price has reduced
		//so that a purchase can be made
		if inp[i] < buyPrice {
			buyPrice = inp[i]
		}

		//check if profit is higher
		if inp[i]-buyPrice > ans {
			ans = inp[i] - buyPrice
		}
	}
	return ans
}
