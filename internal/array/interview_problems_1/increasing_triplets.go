package interviewproblems1

/*
Problem Description
You are given an array A of N elements. Find the number of triplets i,j and k such that i<j<k and A[i]<A[j]<A[k]

Problem Constraints
1 <= N <= 103
1 <= A[i] <= 109

Input Format
First argument A is an array of integers.

Output Format
Return an integer.

Example Input

Input 1:
A = [1, 2, 4, 3]

Input 2:
A = [2, 1, 2, 3]

Example Output
Output 1:
2

Output 2:
1

Example Explanation
For Input 1:
The triplets that satisfy the conditions are [1, 2, 3] and [1, 2, 4].

For Input 2:
The triplet that satisfy the conditions is [1, 2, 3].
*/
func GetTripletCount(A []int) int {
	tripletCount := 0
	//iterate through middle elements
	for i := 1; i < len(A)-1; i++ {
		smallerCount := 0
		//count smaller elements
		for j := i - 1; j >= 0; j-- {
			if A[j] < A[i] {
				smallerCount++
			}
		}

		largerCount := 0
		//count larger elements
		for k := i + 1; k < len(A); k++ {
			if A[k] > A[i] {
				largerCount++
			}
		}
		//count triplets
		tripletCount += smallerCount * largerCount
	}
	return tripletCount
}
