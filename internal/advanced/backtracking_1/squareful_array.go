package backtracking1

import "math"

/*
Problem Description
Given an array of integers A, the array is squareful if for every pair of adjacent elements, their sum is a perfect square.
Find and return the number of permutations of A that are squareful. Two permutations A1 and A2 differ if and only if there is some index i such that A1[i] != A2[i].

Problem Constraints
1 <= length of the array <= 12
1 <= A[i] <= 109

Input Format
The only argument given is the integer array A.

Output Format
Return the number of permutations of A that are squareful.

Example Input
Input 1:
A = [2, 2, 2]

Input 2:
A = [1, 17, 8]

Example Output
Output 1:
1

Output 2:
2

Example Explanation
Explanation 1:
Only permutation is [2, 2, 2], the sum of adjacent element is 4 and 4 and both are perfect square.

Explanation 2:
Permutation are [1, 8, 17] and [17, 8, 1].
*/

func GetSquarefulCount(A []int) int {
	if len(A) < 2 {
		return 0
	}
	hm := make(map[int]int)
	for i := range A {
		hm[A[i]]++
	}

	//create a freq array to save the element with their frequency
	freqArr := make([]info, len(hm))
	i := 0
	for k, v := range hm {
		freqArr[i].value = k
		freqArr[i].freq = v
		i++
	}
	cur := make([]int, len(A))

	return getSquareFul(0, freqArr, cur)
}

func getSquareFul(idx int, freqArr []info, cur []int) int {
	if idx == len(cur) {
		return 1
	}
	ans := 0

	//iterate over freq array
	for i := 0; i < len(freqArr); i++ {

		//if the freq >0, include the element
		if freqArr[i].freq > 0 {
			cur[idx] = freqArr[i].value

			//if the idx is >=1, check the last pairs if they're forming a perfect square
			if (idx == 0) || idx > 0 && isPerfectSquare(cur[idx]+cur[idx-1]) {
				freqArr[i].freq--
				ans += getSquareFul(idx+1, freqArr, cur)
				freqArr[i].freq++
			}
		}
	}
	return ans
}

func isPerfectSquare(inp int) bool {
	sqrt := int(math.Sqrt(float64(inp)))
	return sqrt*sqrt == inp
}
