package dp3

import "math"

/*
Problem Description
Given an array A of positive elements, you have to flip the sign of some of its elements such that the resultant sum of the elements of array should be minimum non-negative(as close to zero as possible).
Return the minimum number of elements whose sign needs to be flipped such that the resultant sum is minimum non-negative.

Problem Constraints
1 <= length of(A) <= 100
Sum of all the elements will not exceed 10,000.

Input Format
First and only argument is an integer array A.

Output Format
Return an integer denoting the minimum number of elements whose sign needs to be flipped.

Example Input
Input 1:
A = [15, 10, 6]

Input 2:
A = [14, 10, 4]

Example Output
Output 1:
1

Output 2:
1

Example Explanation
Explanation 1:
Here, we will flip the sign of 15 and the resultant sum will be 1.

Explanation 2:
Here, we will flip the sign of 14 and the resultant sum will be 0.
Note that flipping the sign of 10 and 4 also gives the resultant sum 0 but flippings there sign are not minimum.
*/
func GetMinFlips(A []int) int {
	N := len(A)
	inp := A
	sum := 0
	for i := range inp {
		sum += inp[i]
	}
	sum = sum >> 1
	memo := make([][]int, N+1)
	for i := range memo {
		memo[i] = make([]int, sum+1)
		//keep the 0th col of every row as 0
		//because to store 0 sum, min items to include is 0
		for j := 1; j <= sum; j++ {
			memo[i][j] = math.MaxInt
		}
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= sum; j++ {
			//if we don't pick the item
			memo[i][j] = memo[i-1][j]

			//if we pick the item, add 1 for the inclusion count
			if j >= inp[i-1] && memo[i-1][j-inp[i-1]] != math.MaxInt {
				memo[i][j] = min(memo[i][j], 1+memo[i-1][j-inp[i-1]])
			}
		}
	}

	col := sum
	for memo[N][col] == math.MaxInt {
		col--
	}
	return memo[N][col]
}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}
