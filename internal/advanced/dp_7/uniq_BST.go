package dp7

/*
Problem Description
Given an integer A, how many structurally unique BST's (binary search trees) exist that can store values 1...A?

Problem Constraints
1 <= A <=18

Input Format
First and only argument is the integer A

Output Format
Return a single integer, the answer to the problem

Example Input
Input 1:
1

Input 2:
2

Example Output
Output 1:
1

Output 2:
2

Example Explanation
Explanation 1:
Only single node tree is possible.

Explanation 2:
2 trees are possible, one rooted at 1 and other rooted at 2.
*/
//recurrence relation: state(x)=sum(state(y-1)*state(x-y)) where y varies from 1 to x
func GetUniqBST(A int) int {
	inp := A
	memo := make([]int, inp+1)
	if A == 1 || A == 0 {
		return 1
	}
	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= inp; i++ {
		ans := 0
		for j := 1; j <= i; j++ {
			ans += memo[j-1] * memo[i-j]
		}
		memo[i] = ans
	}
	return memo[inp]
}
