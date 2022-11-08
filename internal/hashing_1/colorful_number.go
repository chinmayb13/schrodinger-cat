package hashing1

/*
Problem Description
Given a number A, find if it is COLORFUL number or not.
If number A is a COLORFUL number return 1 else, return 0.

What is a COLORFUL Number:

A number can be broken into different contiguous sub-subsequence parts.
Suppose, a number 3245 can be broken into parts like 3 2 4 5 32 24 45 324 245.
And this number is a COLORFUL number, since product of every digit of a contiguous subsequence is different.

Problem Constraints
1 <= A <= 2 * 109

Input Format
The first and only argument is an integer A.

Output Format
Return 1 if integer A is COLORFUL else return 0.

Example Input
Input 1:
A = 23

Input 2:
A = 236

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
Possible Sub-sequences: [2, 3, 23] where
2 -> 2
3 -> 3
23 -> 6  (product of digits)
This number is a COLORFUL number since product of every digit of a sub-sequence are different.

Explanation 2:
Possible Sub-sequences: [2, 3, 6, 23, 36, 236] where
2 -> 2
3 -> 3
6 -> 6
23 -> 6  (product of digits)
36 -> 18  (product of digits)
236 -> 36  (product of digits)
This number is not a COLORFUL number since the product sequence 23  and sequence 6 is same.
*/
func IsColorful(A int) int {
	var digitArray []int
	for A > 0 {
		digitArray = append(digitArray, A%10)
		A = A / 10
	}

	hashSet := make(map[uint64]struct{})
	for l := range digitArray {
		var product uint64 = 1
		for r := l; r < len(digitArray); r++ {
			product *= uint64(digitArray[r])
			if _, ok := hashSet[product]; ok {
				return 0
			} else {
				hashSet[product] = struct{}{}
			}
		}
	}
	return 1
}
