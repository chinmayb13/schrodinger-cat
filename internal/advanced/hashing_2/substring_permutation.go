package hashing2

/*
Problem Description
You are given two strings, A and B, of size N and M, respectively.
You have to find the count of all permutations of A present in B as a substring. You can assume a string will have only lowercase letters.

Problem Constraints
1 <= N < M <= 105

Input Format
Given two arguments, A and B of type String.

Output Format
Return a single integer, i.e., number of permutations of A present in B as a substring.

Example Input
Input 1:
A = "abc"
B = "abcbacabc"

Input 2:
A = "aca"
B = "acaa"

Example Output
Output 1:
5

Output 2:
2

Example Explanation
Explanation 1:
Permutations of A that are present in B as substring are:
 1. abc
 2. cba
 3. bac
 4. cab
 5. abc
    So ans is 5.

Explanation 2:
Permutations of A that are present in B as substring are:
 1. aca
 2. caa
*/
func GetSubstringPermutation(A string, B string) int {
	ans := 0
	var prev, next byte
	var expectedLength, actualLength int
	subStrFreqMap := make(map[byte]int)
	strFreqMap := make(map[byte]int)
	for i := range A {
		subStrFreqMap[A[i]]++
	}

	expectedLength = len(subStrFreqMap)
	//check for first A characters
	for i := 0; i < len(A); i++ {
		strFreqMap[B[i]]++
		if strFreqMap[B[i]] == subStrFreqMap[B[i]] {
			actualLength++
		}
	}

	if actualLength == expectedLength {
		ans++
	}

	//use sliding window, if current window is matching to A, increase the ans count
	l, r := 1, len(A)
	for r < len(B) {
		prev = B[l-1]
		next = B[r]

		if strFreqMap[prev] == subStrFreqMap[prev] {
			actualLength--
		}
		strFreqMap[prev]--

		strFreqMap[next]++
		if strFreqMap[next] == subStrFreqMap[next] {
			actualLength++
		}

		if actualLength == expectedLength {
			ans++
		}
		l++
		r++
	}
	return ans
}
