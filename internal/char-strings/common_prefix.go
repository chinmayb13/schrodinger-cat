package charstrings

/*
Problem Description
Given the array of strings A, you need to find the longest string S, which is the prefix of ALL the strings in the array.
The longest common prefix for a pair of strings S1 and S2 is the longest string S which is the prefix of both S1 and S2.
Example: the longest common prefix of "abcdefgh" and "abcefgh" is "abc".

Problem Constraints
0 <= sum of length of all strings <= 1000000

Input Format
The only argument given is an array of strings A.

Output Format
Return the longest common prefix of all strings in A.

Example Input
Input 1:
A = ["abcdefgh", "aefghijk", "abcefgh"]

Input 2:
A = ["abab", "ab", "abcd"];

Example Output
Output 1:
"a"

Output 2:
"ab"

Example Explanation
Explanation 1:
Longest common prefix of all the strings is "a".

Explanation 2:
Longest common prefix of all the strings is "ab".
*/
func LongestCommonPrefix(A []string) string {
	if len(A) == 0 {
		return ""
	}
	//assume the first string is the common prefix
	prefixString := A[0]
	end := len(prefixString)
	//iterate over the subsequent strings
	for i := 1; i < len(A); i++ {
		j := 0
		//iterate chars within the string
		for j < end && j < len(A[i]) {
			//if the char doesn't match break
			if A[i][j] != prefixString[j] && j < end {
				break
			}
			j++
		}
		//if the new found prefix is of shorter length, then update
		if j < end {
			end = j
		}
	}
	return prefixString[:end]
}
