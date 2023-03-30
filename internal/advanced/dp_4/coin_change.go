package dp4

/*
Problem Description
You are given a set of coins A. In how many ways can you make sum B assuming you have infinite amount of each coin in the set.

NOTE:

Coins in set A will be unique. Expected space complexity of this problem is O(B).
The answer can overflow. So, return the answer % (106 + 7).


Problem Constraints
1 <= A <= 500
1 <= A[i] <= 1000
1 <= B <= 50000

Input Format
First argument is an integer array A representing the set.
Second argument is an integer B.

Output Format
Return an integer denoting the number of ways.

Example Input
Input 1:
A = [1, 2, 3]
B = 4

Input 2:
A = [10]
B = 10

Example Output
Output 1:
4

Output 2:
1

Example Explanation
Explanation 1:
The 4 possible ways are:
{1, 1, 1, 1}
{1, 1, 2}
{2, 2}
{1, 3}

Explanation 2:
There is only 1 way to make sum 10.
*/
func Coinchange2(A []int, B int) int {
	N := len(A)
	var mod int = 1e6 + 7
	cur := make([]int, B+1)
	//state identifier state(idx,x): no of ways of storing a sum x, considering indexes idx to M-1 , where M is B
	//cur[0]=1 since no of ways of storing a sum is always 1
	cur[0] = 1
	next := make([]int, B+1)
	//loop on indexes in reversing order
	for idx := N - 1; idx >= 0; idx-- {
		//loop on sum from 1 to B
		for sum := 1; sum <= B; sum++ {
			cur[sum] = next[sum]
			if sum >= A[idx] {
				cur[sum] = (cur[sum] + cur[sum-A[idx]]) % mod
			}
		}
		next, cur = cur, next
		for i := range cur {
			cur[i] = 0
		}
		cur[0] = 1
	}
	return next[B]
}
