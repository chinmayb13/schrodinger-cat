package stringswithhashmaps

/*
Problem Description
You are given a string A of length N consisting of lowercase alphabets. Find the period of the string.
Period of the string is the minimum value of k (k >= 1), that satisfies A[i] = A[i % k] for all valid i.

Problem Constraints
1 <= N <= 106

Input Format
First and only argument is a string A of length N.

Output Format
Return an integer, denoting the period of the string.

Example Input
Input 1:
A = "abababab"

Input 2:
A = "aaaa"

Example Output
Output 1:
2

Output 2:
1

Example Explanation
Explanation 1:
Period of the string will be 2:
Since, for all i, A[i] = A[i%2].

Explanation 2:
Period of the string will be 1.
*/
func GetStringPeriod(A string) int {
	lps := computeLPSArray(A)
	/*
	   Idea: any string would repeat at its length idx
	   Example: abc would start repeating from 3rd idx (abcabc)

	   In a given string, the last element in LPS array would tell how many elements are matching with the beginning of the array
	   String: abcdab, LPS: [000012]
	   For the string to repeat entirely, we want idx = len(String) = 6
	   2 elements are already repeating, remaining needed = 6 - 2 = 2
	*/
	return len(A) - lps[len(lps)-1]
}

func GetStringPeriodAlt(A string) int {
	ans := -1
	lps := computeLPSArray(A)
	i := len(lps) - 1
	//start from right and look out for non-zero sequence
	//keep seraching until we encounter 1, because thats
	//from where element started repeating
	for i >= 0 && lps[i] > 0 {
		if lps[i] == 1 {
			ans = i
			break
		}
		i--
	}
	if ans == -1 {
		return len(lps)
	}
	return ans
}
