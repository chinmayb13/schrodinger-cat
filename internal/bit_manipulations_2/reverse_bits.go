package bitmanipulations2

/*
Problem Description
Reverse the bits of an 32 bit unsigned integer A.

Problem Constraints
0 <= A <= 232

Input Format
First and only argument of input contains an integer A.

Output Format
Return a single unsigned integer denoting the decimal value of reversed bits.

Example Input
Input 1:
0

Input 2:
3

Example Output
Output 1:
0

Output 2:
3221225472

Example Explanation
Explanation 1:

	00000000000000000000000000000000

=>      00000000000000000000000000000000

Explanation 2:

	00000000000000000000000000000011

=>      11000000000000000000000000000000
*/
func ReverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; num > 0; i++ {
		//for every set bit, set its opposite position's(31-i) bit
		if num&1 > 0 {
			rev |= (1 << (31 - i))
		}
		num = num >> 1
	}
	return rev
}
