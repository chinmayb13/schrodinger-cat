package hashing2

/*
Problem Description
Given a string A, find the length of the longest substring without repeating characters.
Note: Users are expected to solve in O(N) time complexity.

Problem Constraints
1 <= size(A) <= 106
String consists of lowerCase,upperCase characters and digits are also present in the string A.

Input Format
Single Argument representing string A.

Output Format
Return an integer denoting the maximum possible length of substring without repeating characters.

Example Input
Input 1:
A = "abcabcbb"

Input 2:
A = "AaaA"

Example Output
Output 1:
3

Output 2:
2

Example Explanation
Explanation 1:
Substring "abc" is the longest substring without repeating characters in string A.

Explanation 2:
Substring "Aa" or "aA" is the longest substring without repeating characters in string A.
*/
func GetLongestUniqueSubstring(A string) int {
	var maxLength, length int
	l, r := 0, 0
	freqMap := make(map[byte]int)
	for r < len(A) {
		freqMap[A[r]]++
		for freqMap[A[r]] > 1 {
			freqMap[A[l]]--
			l++
		}
		length = r - l + 1
		if length > maxLength {
			maxLength = length
		}
		r++
	}
	return maxLength
}
