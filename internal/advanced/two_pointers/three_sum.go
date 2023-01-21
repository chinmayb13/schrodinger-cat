package twopointers

import (
	"sort"
)

/*
Problem Description
Given an array A of N integers, find three integers in A such that the sum is closest to a given number B. Return the sum of those three integers.
Assume that there will only be one solution.

Problem Constraints
-108 <= B <= 108
1 <= N <= 104
-108 <= A[i] <= 108

Input Format
First argument is an integer array A of size N.
Second argument is an integer B denoting the sum you need to get close to.

Output Format
Return a single integer denoting the sum of three integers which is closest to B.

Example Input
Input 1:
A = [-1, 2, 1, -4]
B = 1

Input 2:
A = [1, 2, 3]
B = 6

Example Output
Output 1:
2

Output 2:
6

Example Explanation
Explanation 1:
The sum that is closest to the target is 2. (-1 + 2 + 1 = 2)

Explanation 2:
Take all elements to get exactly 6.
*/
func ThreeSumClosest(A []int, B int) int {
	if len(A) < 3 {
		return -1
	}
	sort.Ints(A)
	var ans int = 1e9
	sum := 0
	for i := 0; i < len(A)-2; i++ {
		ptr1 := i + 1
		ptr2 := len(A) - 1
		for ptr1 < ptr2 {
			sum = A[i] + A[ptr1] + A[ptr2]
			if abs(sum-B) < abs(ans-B) {
				ans = sum
			}
			if sum > B {
				ptr2--
			} else {
				ptr1++
			}
		}
	}
	return ans
}

func abs(inp int) int {
	if inp < 0 {
		return inp * -1
	}
	return inp
}
