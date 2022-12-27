package combinatorics

import "sort"

/*
Problem Description
Given a string A. Find the rank of the string amongst its permutations sorted lexicographically.
Assume that no characters are repeated.

Note: The answer might not fit in an integer, so return your answer % 1000003

Problem Constraints
1 <= |A| <= 1000

Input Format
First argument is a string A.

Output Format
Return an integer denoting the rank of the given string.

Example Input
Input 1:
A = "acb"

Input 2:
A = "a"

Example Output
Output 1:
2

Output 2:
1

Example Explanation
Explanation 1:
Given A = "acb".
The order permutations with letters 'a', 'c', and 'b' :
abc
acb
bac
bca
cab
cba
So, the rank of A is 2.

Explanation 2:
Given A = "a".
Rank is clearly 1.
*/

func FindRank(A string) int {
	ans := 0
	mod := 1000003
	for i := 0; i < len(A)-1; i++ {
		count := 0
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[i] {
				count++
			}
		}
		ans = (ans + (count*factorial(len(A)-i-1))%mod) % mod
	}
	return (ans + 1) % mod
}

func FindRankAlt(A string) int {
	rank := 1
	charsVisited := 0
	mod := 1000003
	byteArr := []byte(A)
	//sort the string
	sort.Slice(byteArr, func(i, j int) bool {
		return byteArr[i] < byteArr[j]
	})
	for i := 0; i < len(A); i++ {
		position := 0
		for j := 0; j < len(byteArr); j++ {
			//increase the counter, if there are characters present which are lexicographically smaller than current input character
			if byteArr[j] != A[i] && byteArr[j] != '0' {
				position++
			} else if byteArr[j] == A[i] {
				//increase the counter, for how many input charactrers has been visited so far
				charsVisited++
				//increase the rank by position count * all combinations of unvisited elements
				rank = (rank + (position*factorial(len(byteArr)-charsVisited))%mod) % mod
				byteArr[j] = '0'
				break
			}
		}
	}
	return rank
}

func factorial(inp int) int {
	product := 1
	mod := 1000003
	for i := 2; i <= inp; i++ {
		product = (product * i) % mod
	}
	return product
}
