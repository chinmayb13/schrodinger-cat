package stringswithhashmaps

import "strings"

/*
Problem Description
Given two binary strings A and B, count how many cyclic shift of B when taken XOR with A give 0.
NOTE: If there is a string, S0, S1, ... Sn-1 , then it is a cyclic shift is of the form Sk, Sk+1, ... Sn-1, S0, S1, ... Sk-1 where k can be any integer from 0 to N-1.

Problem Constraints
1 ≤ length(A) = length(B) ≤ 105

Input Format
The first argument is a string A.
The second argument is a string B.

Output Format
Return an integer denoting the required answer.

Example Input
Input 1:
A = "1001"
B = "0011"

Input 2:
A = "111"
B = "111"

Example Output
Output 1:
1

Output 2:
3

Example Explanation
Explanation 1:
4 cyclic shifts of B exists: "0011", "0110", "1100", "1001".
There is only one cyclic shift of B i.e. "1001" which has 0 xor with A.

Explanation 2:
All cyclic shifts of B are same as A and give 0 when taken xor with A. So, the ans is 3.
*/

func GetSimilarPatternCount(A string, B string) int {
	var sb strings.Builder
	var ans int
	inputLength := len(B)
	patternLength := len(A)
	sb.WriteString(A)
	sb.WriteString("@")
	sb.WriteString(B)
	sb.WriteString(B[:inputLength-1])
	lps := computeLPSArray(sb.String())
	for i := (patternLength << 1); i < len(lps); i++ {
		if lps[i] == patternLength {
			ans++
		}
	}
	return ans
}

// KMP algorithm to build LPS array
func computeLPSArray(inp string) []int {
	lps := make([]int, len(inp))
	var x int
	lps[0] = 0
	for i := 1; i < len(lps); i++ {
		x = lps[i-1]
		for inp[i] != inp[x] {
			if x == 0 {
				x = -1
				break
			}
			x = lps[x-1]
		}
		lps[i] = x + 1
	}
	return lps
}
