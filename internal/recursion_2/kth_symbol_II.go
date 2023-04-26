package recursion2

/*
Problem Description
On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.
Given row number A and index B, return the Bth indexed symbol in row A. (The values of B are 0-indexed.).

Problem Constraints
1 <= A <= 105
0 <= B <= min(2A - 1 - 1 , 1018)

Input Format
First argument is an integer A.
Second argument is an integer B.

Output Format
Return an integer denoting the Bth indexed symbol in row A.

Example Input
Input 1:
A = 3
B = 0

Input 2:
A = 4
B = 4

Example Output
Output 1:
0

Output 2:
1

Example Explanation
Explanation 1:
Row 1: 0
Row 2: 01
Row 3: 0110

Explanation 2:
Row 1: 0
Row 2: 01
Row 3: 0110
Row 4: 01101001
*/

/*
approach:

	0
	01
	0110
	01101001

If we look at the pattern, element at ith index and element at i/2th index are same when i is even,
while they are changing when the i is odd
*/
func GetKthSymbolII(A int, B int) int {
	if B == 0 {
		return 0
	}
	if A == 1 {
		return 0
	}
	if B&1 == 0 {
		return GetKthSymbolII(A-1, B/2)
	}
	return 1 ^ GetKthSymbolII(A-1, B/2)
}
