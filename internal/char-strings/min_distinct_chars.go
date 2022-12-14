package charstrings

import "sort"

/*
Problem Description
You are given a string A of size N consisting of lowercase alphabets.
You can change at most B characters in the given string to any other lowercase alphabet such that the number of distinct characters in the string is minimized.
Find the minimum number of distinct characters in the resulting string.

Problem Constraints
1 <= N <= 100000
0 <= B < N

Input Format
The first argument is a string A.
The second argument is an integer B.

Output Format
Return an integer denoting the minimum number of distinct characters in the string.

Example Input
A = "abcabbccd"
B = 3

Example Output
2

Example Explanation
We can change both 'a' and one 'd' into 'b'.So the new string becomes "bbcbbbccb".
So the minimum number of distinct character will be 2.
*/
func GetMinDistinctChars(A string, B int) int {
	charArr := make([]int, 26)
	uniqueCount := 0
	for i := range A {
		if charArr[A[i]-'a'] == 0 {
			uniqueCount++
		}
		charArr[A[i]-'a']++
	}
	sort.Ints(charArr)

	for i := range charArr {
		frequency := charArr[i]
		if 0 < frequency && frequency <= B && uniqueCount > 1 {
			B -= charArr[i]
			uniqueCount--
		}

	}
	return uniqueCount
}

func GetMinDistinctCharsAlt(A string, B int) int {
	charArr := make([]int, 26)
	count := 0
	for i := range A {
		charArr[A[i]-'a']++

	}
	sort.Ints(charArr)
	for i := range charArr {
		frequency := charArr[i]
		if frequency <= B && B > 0 {
			B -= frequency
		} else if frequency > 0 {
			count++
		}

	}
	if count == 0 {
		return 1
	}
	return count
}
