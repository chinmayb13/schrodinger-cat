package gcd

/*
Problem Description
Given an integer array A of size N. You have to delete one element such that the GCD(Greatest common divisor) of the remaining array is maximum.
Find the maximum value of GCD.

Problem Constraints
2 <= N <= 105
1 <= A[i] <= 109

Input Format
First argument is an integer array A.

Output Format
Return an integer denoting the maximum value of GCD.

Example Input
Input 1:
A = [12, 15, 18]

Input 2:
A = [5, 15, 30]

Example Output
Output 1:
6

Output 2:
15

# Example Explanation

Explanation 1:
If you delete 12, gcd will be 3.
If you delete 15, gcd will be 6.
If you delete 18, gcd will 3.
Maximum vallue of gcd is 6.

Explanation 2:
If you delete 5, gcd will be 15.
If you delete 15, gcd will be 5.
If you delete 30, gcd will be 5.
*/
func GetMaxGCD(A []int) int {
	prefixG := buildPrefixGCD(A)
	maxGCD := 0
	gcd := 0
	for i := len(A) - 1; i >= 0; i-- {
		localGCD := getGCD(gcd, prefixG[i])
		if localGCD > maxGCD {
			maxGCD = localGCD
		}
		gcd = getGCD(gcd, A[i])
	}
	return maxGCD
}

func buildPrefixGCD(inp []int) []int {
	var prefixG []int
	gcd := 0
	for i := range inp {
		prefixG = append(prefixG, gcd)
		gcd = getGCD(gcd, inp[i])
	}
	return prefixG
}
