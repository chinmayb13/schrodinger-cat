package charstrings

/*
Problem Description
You are given two lowercase strings A and B each of length N. Return 1 if they are anagrams to each other and 0 if not.

Note : Two strings A and B are called anagrams to each other if A can be formed after rearranging the letters of B.

Problem Constraints
1 <= N <= 105
A and B are lowercase strings

Input Format
Both arguments A and B are a string.

Output Format
Return 1 if they are anagrams and 0 if not

Example Input
Input 1:
A = "cat"
B = "bat"

Input 2:
A = "secure"
B = "rescue"

Example Output
Output 1:
0

Output 2:
1

Example Explanation
For Input 1:
The words cannot be rearranged to form the same word. So, they are not anagrams.

For Input 2:
They are an anagram.
*/
func IsAnagram(A string, B string) int {
	if len(A) != len(B) {
		return 0
	}

	baseElem := 'a'
	countArr_a := make([]int, 26)
	countArr_b := make([]int, 26)
	for i := 0; i < len(A); i++ {
		countArr_a[A[i]-byte(baseElem)]++
		countArr_b[B[i]-byte(baseElem)]++
	}

	for i := 0; i < 26; i++ {
		if countArr_a[i] != countArr_b[i] {
			return 0
		}
	}
	return 1
}
