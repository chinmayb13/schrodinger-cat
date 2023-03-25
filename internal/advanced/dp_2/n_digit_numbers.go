package dp2

/*
Problem Description
Find out the number of A digit positive numbers, whose digits on being added equals to a given number B.
Note that a valid number starts from digits 1-9 except the number 0 itself. i.e. leading zeroes are not allowed.
Since the answer can be large, output answer modulo 1000000007

Problem Constraints
1 <= A <= 1000
1 <= B <= 10000

Input Format
First argument is the integer A
Second argument is the integer B

Output Format
Return a single integer, the answer to the problem

Example Input
Input 1:
A = 2
B = 4

Input 2:
A = 1
B = 3

Example Output
Output 1:
4

Output 2:
1

Example Explanation
Explanation 1:
Valid numbers are {22, 31, 13, 40}
Hence output 4.

Explanation 2:
Only valid number is 3
*/
func GetNumberCount(A int, B int) int {
	var mod int = 1e9 + 7
	sum := B
	digits := A
	memo := make([][]int, A+1)
	for i := range memo {
		memo[i] = make([]int, B+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	ans := 0
	//first digit can extend from 1 to min(9,sum)
	for i := 1; i <= min(B, 9); i++ {
		ans = (ans + countNumbers(sum-i, digits-1, memo, mod)) % mod
	}
	return ans
}

func countNumbers(sum int, digits int, memo [][]int, mod int) int {
	//base condition
	//if remaining digits is 0, check for sum
	if digits == 0 {
		if sum == 0 {
			return 1
		}
		return 0
	}

	if memo[digits][sum] == -1 {
		ans := 0
		//second digit onwards can extend from 0 to min(sum,9)
		for i := 0; i <= min(sum, 9); i++ {
			ans = (ans + countNumbers(sum-i, digits-1, memo, mod)) % mod
		}
		memo[digits][sum] = ans
	}
	return memo[digits][sum]

}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}
