package bitmanipulations2

/*
Problem Description
Given an integer A. Unset B bits from the right of A in binary.

For eg:-
A = 93, B = 4
A in binary = 1011101
A should become = 1010000 = 80. Therefore return 80.

Problem Constraints
1 <= A <= 1018
1 <= B <= 60

Input Format
The first argument is a single integer A.
The second argument is a single integer B.

Output Format
Return the number with B unset bits from the right.

Example Input
Input 1:-
A = 25
B = 3

Input 2:-
A = 37
B = 3

Example Output
Output 1:-
24

Output 2:-
32

Example Explanation
Explanation 1:-
A = 11001 to 11000

Explantion 2:-
A = 100101 to 100000
*/
func UnsetRightBits(A int64, B int) int64 {
	var shift int64 = 1
	for i := 1; i < B; i++ {
		shift = shift | (1 << i)
	}
	return A & (^(shift))
}
