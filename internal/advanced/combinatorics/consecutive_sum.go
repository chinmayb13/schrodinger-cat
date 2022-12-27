package combinatorics

/*
Problem Description
Given a positive integer A.

Return the number of ways it can be written as a sum of consecutive positive integers.

Problem Constraints
1 <= A <= 109

Input Format
The first and the only argument of input contains an integer, A.

Output Format
Return an integer, representing the answer as described in the problem statement.

Example Input
Input 1:
A = 5

Input 2:
A = 15

Example Output
Output 1:
2

Output 2:
4

Example Explanation
Explanation 1:
The 2 ways are:
 => [5]
 => [2, 3]

Explanation 2:
The 4 ways are:
 => [15]
 => [7, 8]
 => [4, 5, 6]
 => [1, 2, 3, 4, 5]
*/

/*
Explanation:
Assume x + x+1 + x+2.... x+k = A
=> x*(k+1) + k*(k+1)/2 = A
=> x*(k+1) = A - k*(k+1)/2
=> x = (A - k*(k+1)/2)/(k+1)
*/
func GetWaysForConsecutiveSum(A int) int {
	count := 0
	var i int64 = 1
	for {
		var sum int64 = i * (i + 1) / 2
		if sum >= int64(A) {
			break
		}
		if (int64(A)-sum)%(i+1) == 0 {
			count++
		}
		i++
	}
	return count + 1
}
