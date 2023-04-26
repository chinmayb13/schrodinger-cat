package dp1

/*
Problem Description
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26

Given an encoded message denoted by string A containing digits, determine the total number of ways to decode it modulo 109 + 7.

Problem Constraints
1 <= length(A) <= 105

Input Format
The first and the only argument is a string A.

Output Format
Return an integer, representing the number of ways to decode the string modulo 109 + 7.

Example Input
Input 1:
A = "12"

Input 2:
A = "8"

Example Output
Output 1:
2

Output 2:
1

Example Explanation
Explanation 1:
Given encoded message "12", it could be decoded as "AB" (1, 2) or "L" (12).
The number of ways decoding "12" is 2.

Explanation 2:
Given encoded message "8", it could be decoded as only "H" (8).
The number of ways decoding "8" is 1.
*/
func NumDecodings(A string) int {
	memo := make([]int, len(A))
	for i := range memo {
		memo[i] = -1
	}
	return countDecoding(0, A, 1e9+7, memo)
}

func countDecoding(idx int, inp string, mod int, memo []int) int {
	//if end of string, add 1 to the answer
	if idx == len(inp) {
		return 1
	}

	//if current element is 0, return 0
	if inp[idx] == '0' {
		return 0
	}

	//if memo entry is not present
	if memo[idx] == -1 {

		// jump to next character
		memo[idx] = countDecoding(idx+1, inp, mod, memo)

		// if the current digit pair can also be mapped to a character,
		// jump to the next character
		if idx+1 < len(inp) && inp[idx:idx+2] <= "26" {
			memo[idx] = (memo[idx] + countDecoding(idx+2, inp, mod, memo)) % mod
		}
	}
	return memo[idx]
}
