package dp4

/*
Problem Description
Given a rod of length N units and an array A of size N denotes prices that contains prices of all pieces of size 1 to N.

Find and return the maximum value that can be obtained by cutting up the rod and selling the pieces.

Problem Constraints
1 <= N <= 1000
0 <= A[i] <= 106

Input Format
First and only argument is an integer array A of size N.

Output Format
Return an integer denoting the maximum value that can be obtained by cutting up the rod and selling the pieces.

Example Input
Input 1:
A = [3, 4, 1, 6, 2]

Input 2:
A = [1, 5, 2, 5, 6]

Example Output
Output 1:
15

Output 2:
11

Example Explanation
Explanation 1:
Cut the rod of length 5 into 5 rods of length (1, 1, 1, 1, 1) and sell them for (3 + 3 + 3 + 3 + 3) = 15.

Explanation 2:
Cut the rod of length 5 into 3 rods of length (2, 2, 1) and sell them for (5 + 5 + 1) = 11.
*/
func GetMaxProfit(A []int) int {
	inp := A
	N := len(inp)
	memo := make([]int, N+1)

	//for items from 1 to N
	for i := 1; i <= N; i++ {
		//take all possible combinations and take max
		for k := 1; k <= i; k++ {
			memo[i] = max(memo[i], memo[i-k]+inp[k-1])
		}
	}
	return memo[N]
}

func getMaxValueRecursive(length int, inp []int, memo []int) int {
	if length == 0 {
		return 0
	}
	ans := 0
	if memo[length] == -1 {
		for i := 0; i < length; i++ {
			val := getMaxValueRecursive(length-i-1, inp, memo)
			ans = max(ans, inp[i]+val)
		}
		memo[length] = ans
	}
	return memo[length]
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}
