package stringspatternmatching

/*
Problem Description
Given two string A and B of length N and M respectively consisting of lowercase letters. Find the number of occurrences of B in A.

Problem Constraints
1 <= M <= N <= 105

Input Format
Two argument A and B are strings

Output Format
Return an integer denoting the number of occurrences of B in A

Example Input
Input 1:
A = "acbac"
B = "ac"

Input 2:
A = "aaaa"
B = "aa"

Example Output
Output 1:
2

Output 2:
3

Example Explanation
For Input 1:
The string "ac" occurs twice in "acbac".

For Input 2:
The string "aa" occurs thrice in "aaaa".
*/
func GetOccurences(A string, B string) int {
	ans := 0
	target := len(B)
	lps := computeLPSArray(B + "@" + A)
	for i := target << 1; i < len(lps); i++ {
		if lps[i] == target {
			ans++
		}
	}
	return ans
}

func computeLPSArray(inp string) []int {
	arrLength := len(inp)
	lps := make([]int, arrLength)
	len := 0
	lps[0] = 0
	i := 1
	for i < arrLength {
		if inp[i] == inp[len] {
			len++
			lps[i] = len
			i++
		} else if len != 0 {
			len = lps[len-1]
		} else {
			lps[i] = 0
			i++
		}
	}
	return lps
}
