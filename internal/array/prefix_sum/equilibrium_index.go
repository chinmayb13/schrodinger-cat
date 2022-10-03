package prefixsum

/*
You are given an array A of integers of size N.

Your task is to find the equilibrium index of the given array

The equilibrium index of an array is an index such that the sum of elements at lower indexes is equal to the sum of elements at higher indexes.

NOTE:

Array indexing starts from 0.
If there is no equilibrium index then return -1.
If there are more than one equilibrium indexes then return the minimum index.

Problem Constraints
1 <= N <= 105
-105 <= A[i] <= 105


Input Format
First arugment is an array A .


Output Format
Return the equilibrium index of the given array. If no such index is found then return -1.


Example Input
Input 1:
A=[-7, 1, 5, 2, -4, 3, 0]

Input 2:
A=[1,2,3]


Example Output
Output 1:
3

Output 2:
-1


Example Explanation
Explanation 1:
3 is an equilibrium index, because:
A[0] + A[1] + A[2] = A[4] + A[5] + A[6]
Explanation 1:

There is no such index.
*/

func GetEquilIndex(A []int) int {
	var prefixSum []int64

	//create prefix sum array
	for i := range A {
		if i == 0 {
			prefixSum = append(prefixSum, int64(A[i]))
			continue
		}
		prefixSum = append(prefixSum, prefixSum[i-1]+int64(A[i]))
	}

	for i := range prefixSum {
		var leftSum, rightSum int64
		if i == 0 {
			//nothing on the left
			leftSum = 0
		} else {
			//sum until the index element
			leftSum = prefixSum[i-1]
		}

		//sum[L,R]=PS[R] - PS[L-1]
		rightSum = prefixSum[len(A)-1] - prefixSum[i]
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
