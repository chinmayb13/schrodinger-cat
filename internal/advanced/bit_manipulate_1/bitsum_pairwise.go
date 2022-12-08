package bitmanipulate1

/*
Problem Description
We define f(X, Y) as the number of different corresponding bits in the binary representation of X and Y.
For example, f(2, 7) = 2, since the binary representation of 2 and 7 are 010 and 111, respectively. The first and the third bit differ, so f(2, 7) = 2.
You are given an array of N positive integers, A1, A2,..., AN. Find sum of f(Ai, Aj) for all pairs (i, j) such that 1 ≤ i, j ≤ N. Return the answer modulo 109+7.

Problem Constraints
1 <= N <= 105
1 <= A[i] <= 231 - 1

Input Format
The first and only argument of input contains a single integer array A.

Output Format
Return a single integer denoting the sum.

Example Input
Input 1:
A = [1, 3, 5]

Input 2:
A = [2, 3]

Example Output
Ouptut 1:
8

Output 2:
2

Example Explanation
Explanation 1:
f(1, 1) + f(1, 3) + f(1, 5) + f(3, 1) + f(3, 3) + f(3, 5) + f(5, 1) + f(5, 3) + f(5, 5)
= 0 + 1 + 1 + 1 + 0 + 2 + 1 + 2 + 0 = 8

Explanation 2:
f(2, 2) + f(2, 3) + f(3, 2) + f(3, 3) = 0 + 1 + 1 + 0 = 2
*/
func BitDifferenceSum(A []int) int {
	sum := 0
	max := 0
	var mod int = 1e9 + 7
	for i := range A {
		if A[i] > max {
			max = A[i]
		}
	}
	//identify the number of bits to iterate
	bitCount := getBitCount(max)
	for i := 0; i < bitCount; i++ {
		setBitCount := 0
		//for each bit position, count no of set bits
		for j := range A {
			setBitCount += ((A[j] >> i) & 1)
		}
		//count the pairs of 1,0
		sum += (setBitCount * (len(A) - setBitCount)) % mod
	}
	//multiplied because 1,0 should also be calculated once again as 0,1
	return (2 * sum) % mod
}

func getBitCount(A int) int {
	numOfBits := 0
	for A > 0 {
		numOfBits++
		A = A >> 1
	}
	return numOfBits
}
