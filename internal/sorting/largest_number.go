package sorting

import (
	"sort"
	"strconv"
	"strings"
)

/*
Problem Description
Given an array A of non-negative integers, arrange them such that they form the largest number.
Note: The result may be very large, so you need to return a string instead of an integer.

Problem Constraints
1 <= len(A) <= 100000
0 <= A[i] <= 2*109

Input Format
The first argument is an array of integers.

Output Format
Return a string representing the largest number.

Example Input
Input 1:
A = [3, 30, 34, 5, 9]

Input 2:
A = [2, 3, 9, 0]

Example Output
Output 1:
"9534330"

Output 2:
"9320"

Example Explanation
Explanation 1:
Reorder the numbers to [9, 5, 34, 3, 30] to form the largest number.

Explanation 2:
Reorder the numbers to [9, 3, 2, 0] to form the largest number 9320.
*/

func GetLargestNumber(A []int) string {
	sort.Slice(A, func(i, j int) bool {
		exponentI := getExponent(A[i])
		exponentJ := getExponent(A[j])
		num1 := A[i]*exponentJ + A[j]
		num2 := A[j]*exponentI + A[i]
		return num1 > num2
	})

	var largeNumber strings.Builder
	for i := range A {
		if i == 0 && A[i] == 0 {
			return "0"
		}
		largeNumber.WriteString(strconv.Itoa(A[i]))
	}
	return largeNumber.String()
}

func getExponent(inp int) int {
	power := 10
	for (inp / power) > 0 {
		power *= 10
	}
	return power
}
