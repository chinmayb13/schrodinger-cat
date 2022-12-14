package moduloarithmetic

/*
Problem Description
Given an array of integers A, calculate the sum of A [ i ] % A [ j ] for all possible i, j pairs. Return sum % (109 + 7) as an output.

Problem Constraints
1 <= length of the array A <= 105
1 <= A[i] <= 103

Input Format
The only argument given is the integer array A.

Output Format
Return a single integer denoting sum % (109 + 7).

Example Input
Input 1:
A = [1, 2, 3]

Input 2:
A = [17, 100, 11]


Example Output
Output 1:
5

Output 2:
61


Example Explanation
Explanation 1:
(1 % 1) + (1 % 2) + (1 % 3) + (2 % 1) + (2 % 2) + (2 % 3) + (3 % 1) + (3 % 2) + (3 % 3) = 5
*/
/*
Approach:
Create an auxiliary array, for every input array element, increase the counter at the corresponding index in auxiliary array
Now find the every pair, calculate their mod and multiply it with the frequency of numerator value
*/
func GetModSum(A []int) int {
	var mod int = 1e9 + 7
	modSum := 0
	numberList := make([]int, 1001)
	for i := range A {
		numberList[A[i]]++
	}

	for i := range A {
		for j := range numberList {
			if numberList[j] > 0 {
				modSum = (modSum + (j%A[i])*(numberList[j])) % mod
			}

		}
	}
	return modSum
}
