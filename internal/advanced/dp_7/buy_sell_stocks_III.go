package dp7

import "math"

/*
Say you have an array, A, for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete at most 2 transactions.
Return the maximum possible profit.

Note: You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

Input Format:
The first and the only argument is an integer array, A.

Output Format:
Return an integer, representing the maximum possible profit.

Constraints:
1 <= length(A) <= 7e5
1 <= A[i] <= 1e7

Examples:
Input 1:
A = [1, 2, 1, 2]

Output 1:
2

Explanation 1:
Day 0 : Buy
Day 1 : Sell
Day 2 : Buy
Day 3 : Sell

Input 2:
A = [7, 2, 4, 8, 7]

Output 2:
6

Explanation 2:
Day 1 : Buy
Day 3 : Sell
*/
func GetMaxProfitof2Txn(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	buy1, buy2 := math.MinInt, math.MinInt
	sell1, sell2 := 0, 0
	for i := range A {
		//get min A[i] value
		buy1 = max(buy1, -A[i])

		//find profit1
		sell1 = max(sell1, A[i]+buy1)

		//check if buy2 can be made
		buy2 = max(buy2, sell1-A[i])

		//find profit2
		sell2 = max(sell2, A[i]+buy2)
	}
	return sell2
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}
