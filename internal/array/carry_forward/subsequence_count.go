package carryforward

/*
Problem Description
You have given a string A having Uppercase English letters.

You have to find how many times subsequence "AG" is there in the given string.

NOTE: Return the answer modulo 109 + 7 as the answer can be very large.

Problem Constraints
1 <= length(A) <= 105



Input Format
First and only argument is a string A.



Output Format
Return an integer denoting the answer.



Example Input
Input 1:
A = "ABCGAG"

Input 2:
A = "GAB"


Example Output
Output 1:
3

Output 2:
0


Example Explanation

Explanation 1:
Subsequence "AG" is 3 times in given string

Explanation 2:
There is no subsequence "AG" in the given string.

*/

func GetSubSeqCount(A string) int {
	//number of occurences of character G
	countG := 0

	//Total subsequence count
	subSeqCount := 0
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 'G' {
			//increase character G count
			countG++
			//the count would be as many G characters present when A was encountered + previous subsequence count
		} else if A[i] == 'A' {
			subSeqCount += countG
		}
	}
	return subSeqCount % (1e9 + 7)
}
