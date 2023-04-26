package slidingwindow

/*
Problem Description
Given an array A of length N. Also given are integers B and C.
Return 1 if there exists a subarray with length B having sum C and 0 otherwise

Problem Constraints
1 <= N <= 105
1 <= A[i] <= 104
1 <= B <= N
1 <= C <= 109

Input Format
First argument A is an array of integers.

# The remaining arguments B and C are integers

Output Format
Return 1 if such a subarray exist and 0 otherwise

Example Input
Input 1:
A = [4, 3, 2, 6]
B = 2
C = 5

Input 2:
A = [4, 2, 2]
B = 2
C = 8

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
The subarray [3, 2] is of length 2 and sum 5.
Explanation 2:
There are no such subarray.
*/
func SubArraySumExists(A []int, B int, C int) int {
	sum := 0
	//Calculate for first B elements
	for i := 0; i < B; i++ {
		sum += A[i]
	}

	if sum == C {
		return 1
	}

	//Start from left index as 1 and right index as B
	for l, r := 1, B; r < len(A); l, r = l+1, r+1 {
		//remove element from the left and add from the right
		sum = sum - A[l-1] + A[r]
		if sum == C {
			return 1
		}
	}
	return 0
}
