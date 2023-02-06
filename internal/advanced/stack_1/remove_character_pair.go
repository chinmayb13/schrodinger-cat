package stack1

import "strings"

/*
Problem Description
You are given a string A.

An operation on the string is defined as follows:
Remove the first occurrence of the same consecutive characters. eg for a string "abbcd", the first occurrence of same consecutive characters is "bb".
Therefore the string after this operation will be "acd".

Keep performing this operation on the string until there are no more occurrences of the same consecutive characters and return the final string.

Problem Constraints
1 <= |A| <= 100000

Input Format
First and only argument is string A.

Output Format
Return the final string.

Example Input
Input 1:
A = abccbc

Input 2:
A = ab

Example Output
Output 1:
"ac"

Output 2:
"ab"

Example Explanation
Explanation 1:
Remove the first occurrence of same consecutive characters. eg for a string "abbc",
the first occurrence of same consecutive characters is "bb".
Therefore the string after this operation will be "ac".

Explanation 2:
No removals are to be done.
*/
func GetCleanString(A string) string {
	var ans strings.Builder
	s := newStack()
	for i := range A {
		//if second occurence, pop the element
		if A[i] == s.topByte() {
			s.pop()
			//else, push the element
		} else {
			s.push(A[i])
		}
	}
	s1 := newStack()
	for s.size() > 0 {
		s1.push(s.topByte())
		s.pop()
	}

	for s1.size() > 0 {
		ans.WriteByte(s1.topByte())
		s1.pop()
	}
	return ans.String()
}
