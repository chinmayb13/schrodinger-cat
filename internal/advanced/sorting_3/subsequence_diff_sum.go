package sorting3

import "sort"

/*
Problem Description
Given an integer array, A of size N.
You have to find all possible non-empty subsequences of the array of numbers and then, for each subsequence, find the difference between the largest and smallest numbers in that subsequence. Then add up all the differences to get the number.
As the number may be large, output the number modulo 1e9 + 7 (1000000007).

NOTE: Subsequence can be non-contiguous.

Problem Constraints
1 <= N <= 10000
1<= A[i] <=1000

Input Format
First argument is an integer array A.

Output Format
Return an integer denoting the output.

Example Input
Input 1:
A = [1, 2]

Input 2:
A = [1]

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
All possible non-empty subsets are:
[1]    largest-smallest = 1 - 1 = 0
[2]    largest-smallest = 2 - 2 = 0
[1 2]  largest-smallest = 2 - 1 = 1
Sum of the differences = 0 + 0 + 1 = 1
So, the resultant number is 1

Explanation 2:
Only 1 subsequence of 1 element is formed.
*/
func GetSubsequenceDiffSum(A []int) int {
	var mod int = 1e9 + 7
	sort.Ints(A)
	lastIdx := len(A) - 1
	var totMax, totMin int
	for i := range A {
		//the current element will be maximum for subsequences 2^(no of elements before the current element in the sorted array)
		totMax = (totMax + (A[i]*calcPower(2, i))%mod) % mod
		//the current element will be minimum for subsequences 2^(no of elements after the current element in the sorted array)
		totMin = (totMin + (A[i]*calcPower(2, lastIdx-i))%mod) % mod
	}
	return (totMax - totMin + mod) % mod
}

func calcPower(inp, pow int) int {
	var mod int = 1e9 + 7
	ans := 1
	for pow > 0 {
		if (pow & 1) > 0 {
			ans = (ans * inp) % mod
			pow--
		} else {
			inp = (inp * inp) % mod
			pow = pow >> 1
		}
	}
	return ans
}
