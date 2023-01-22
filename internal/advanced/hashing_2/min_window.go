package hashing2

/*
Problem Description
Given a string A and a string B, find the window with minimum length in A, which will contain all the characters in B in linear time complexity.
Note that when the count of a character c in B is x, then the count of c in the minimum window in A should be at least x.

Note:

If there is no such window in A that covers all characters in B, return the empty string.
If there are multiple such windows, return the first occurring minimum window ( with minimum start index and length )

Problem Constraints
1 <= size(A), size(B) <= 106

Input Format
The first argument is a string A.
The second argument is a string B.

Output Format
Return a string denoting the minimum window.

Example Input
Input 1:
A = "ADOBECODEBANC"
B = "ABC"

Input 2:
A = "Aa91b"
B = "ab"

Example Output
Output 1:
"BANC"

Output 2:
"a91b"

Example Explanation
Explanation 1:
"BANC" is a substring of A which contains all characters of B.

Explanation 2:
"a91b" is the substring of A which contains all characters of B.
*/
func GetMinWindow(A string, B string) string {
	var minLength int = len(A) + 1
	var length, beg int
	strFreq := make(map[byte]int)
	subStrFreq := make(map[byte]int)
	//store the freq of each character in B string
	for i := range B {
		subStrFreq[B[i]]++
	}
	//store the count of unique characters in B string
	expectedFreq := len(subStrFreq)
	actualFreq := 0
	l, r := 0, 0
	for r < len(A) {
		strFreq[A[r]]++
		//if the current character is repeated exactly same time in B string, increase unique character count
		if freq, ok := subStrFreq[A[r]]; ok && freq == strFreq[A[r]] {
			actualFreq++
		}

		//once all characters are found with their respective frequency, shrink the substring from left
		for l <= r && actualFreq == expectedFreq {
			length = r - l + 1
			//save the ans before shrinking
			if length < minLength {
				minLength = length
				beg = l
			}
			strFreq[A[l]]--
			if strFreq[A[l]] < subStrFreq[A[l]] {
				actualFreq--
			}
			l++
		}
		r++
	}
	if minLength <= len(A) {
		return A[beg : beg+minLength]
	}
	return ""
}
