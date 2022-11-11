package hashing2

/*
Problem Description
Surprisingly, in an alien language, they also use English lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.
Given an array of words A of size N written in the alien language, and the order of the alphabet denoted by string B of size 26, return 1 if and only if the given words are sorted lexicographically in this alien language else, return 0.

Problem Constraints
1 <= N, length of each word <= 105
Sum of the length of all words <= 2 * 106

Input Format
The first argument is a string array A of size N.
The second argument is a string B of size 26, denoting the order.

Output Format
Return 1 if and only if the given words are sorted lexicographically in this alien language else, return 0.

Example Input
Input 1:
A = ["hello", "scaler", "interviewbit"]
B = "adhbcfegskjlponmirqtxwuvzy"

Input 2:
A = ["fine", "none", "no"]
B = "qwertyuiopasdfghjklzxcvbnm"

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
The order shown in string B is: h < s < i for the given words. So return 1.

Explanation 2:
"none" should be present after "no". Return 0.
*/

//Approach: Take two strings at a time and compare them character by character
func IsDictionary(A []string, B string) int {
	hashMap := make(map[byte]int)
	//create priority deciding map for characters as given in the input
	for i := range B {
		hashMap[B[i]] = i
	}

	for i := 0; i < len(A)-1; i++ {
		charIndex := 0
		//loop until either one of the strings is exhausted or an index is encountered at which the characters are different
		for (charIndex < len(A[i])) && (charIndex < len(A[i+1])) && (hashMap[A[i][charIndex]] == hashMap[A[i+1][charIndex]]) {
			charIndex++
		}
		//if the successor string is exhausted, then it is not sorted (since 'no' comes before 'none')
		if charIndex >= len(A[i+1]) {
			return 0
			//if the predecessor string is exhausted, then move on to next iteration
		} else if charIndex >= len(A[i]) {
			continue
		}
		//if the character present at charIndex of ith string has lower priority than that
		//present at charIndex of (i+1)th string, then it is not sorted
		if hashMap[A[i][charIndex]] > hashMap[A[i+1][charIndex]] {
			return 0
		}
	}
	return 1
}
