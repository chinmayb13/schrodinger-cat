package backtracking1

import "strings"

/*
Problem Description
Given a digit string A, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below.

The digit 0 maps to 0 itself. The digit 1 maps to 1 itself.
NOTE: Make sure the returned strings are lexicographically sorted.

Problem Constraints
1 <= |A| <= 10

Input Format
The only argument is a digit string A.

Output Format
Return a string array denoting the possible letter combinations.

Example Input
Input 1:
A = "23"

Input 2:
A = "012"

Example Output
Output 1:
["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]

Output 2:
["01a", "01b", "01c"]

Example Explanation
Explanation 1:
There are 9 possible letter combinations.

Explanation 2:
Only 3 possible letter combinations.
*/
func LetterCombinations(A string) []string {
	hm := map[byte][]byte{
		'0': {'0'},
		'1': {'1'},
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	cur := make([]byte, len(A))
	var ans []string
	getCombination(A, 0, hm, cur, &ans)
	return ans
}

func getCombination(inp string, idx int, hm map[byte][]byte, cur []byte, ans *[]string) {
	if idx == len(inp) {
		*ans = append(*ans, toString(cur))
		return
	}

	arr := hm[inp[idx]]
	for i := range arr {
		cur[idx] = arr[i]
		getCombination(inp, idx+1, hm, cur, ans)
	}
}

func toString(inp []byte) string {
	var sb strings.Builder
	for i := range inp {
		sb.WriteByte(inp[i])
	}
	return sb.String()
}
