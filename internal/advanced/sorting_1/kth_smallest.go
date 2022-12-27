package sorting1

/*
Problem Description
Find the Bth smallest element in given array A .
NOTE: Users should try to solve it in less than equal to B swaps.

Problem Constraints
1 <= |A| <= 100000
1 <= B <= min(|A|, 500)
1 <= A[i] <= 109

Input Format
The first argument is an integer array A.
The second argument is integer B.

Output Format
Return the Bth smallest element in given array.

Example Input
Input 1:
A = [2, 1, 4, 3, 2]
B = 3

Input 2:
A = [1, 2]
B = 2

Example Output
Output 1:
2

Output 2:
2

Example Explanation
Explanation 1:
3rd element after sorting is 2.

Explanation 2:
2nd element after sorting is 2.
*/

//Approach: Use selection sort algorithm to pick Bth smallest element
func Kthsmallest(A []int, B int) int {
	for i := 0; i < len(A)-1; i++ {
		idx := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[idx] {
				idx = j
			}
		}
		A[idx], A[i] = A[i], A[idx]
		if i+1 == B {
			return A[i]
		}
	}
	return A[len(A)-1]
}
