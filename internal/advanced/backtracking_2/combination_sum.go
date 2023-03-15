package backtracking2

import "sort"

/*
Problem Description
Given an array of candidate numbers A and a target number B, find all unique combinations in A where the candidate numbers sums to B.
The same repeated number may be chosen from A unlimited number of times.

Note:
1) All numbers (including target) will be positive integers.
2) Elements in a combination (a1, a2, … , ak) must be in non-descending order. (ie, a1 ≤ a2 ≤ … ≤ ak).
3) The combinations themselves must be sorted in ascending order.
4) CombinationA > CombinationB iff (a1 > b1) OR (a1 = b1 AND a2 > b2) OR ... (a1 = b1 AND a2 = b2 AND ... ai = bi AND ai+1 > bi+1)
5) The solution set must not contain duplicate combinations.

Problem Constraints
1 <= |A| <= 20
1 <= A[i] <= 50
1 <= B <= 500

Input Format
The first argument is an integer array A.
The second argument is integer B.

Output Format
Return a vector of all combinations that sum up to B.

Example Input
Input 1:
A = [2, 3]
B = 2

Input 2:
A = [2, 3, 6, 7]
B = 7

Example Output
Output 1:
[ [2] ]

Output 2:
[ [2, 2, 3] , [7] ]

Example Explanation
Explanation 1:
All possible combinations are listed.

Explanation 2:
All possible combinations are listed.
*/
func CombinationSum(A []int, B int) [][]int {
	//sort the array
	sort.Ints(A)
	var ans [][]int
	var cur []int
	getCombinations(B, 0, cur, A, &ans)
	return ans
}

func getCombinations(sum int, lastValue int, cur []int, inp []int, ans *[][]int) {
	if sum == 0 {
		*ans = append(*ans, copyIntArr(cur))
		return
	}
	if sum < 0 {
		return
	}

	for i := 0; i < len(inp); i++ {
		//since the array itself should be non-descending order
		if inp[i] >= lastValue {
			cur = append(cur, inp[i])
			getCombinations(sum-inp[i], inp[i], cur, inp, ans)
			cur = cur[:len(cur)-1]

			//to avoid duplicacy, jump to to the next non repeating element
			for i+1 < len(inp) && inp[i] == inp[i+1] {
				i++
			}
		}
	}
}
