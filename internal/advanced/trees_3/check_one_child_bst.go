package trees3

import "math"

/*
Problem Description

Given preorder traversal of a binary tree, check if it is possible that it is also a preorder traversal of a Binary Search Tree (BST), where each internal node (non-leaf nodes) have exactly one child.

Problem Constraints
1 <= number of nodes <= 100000

Input Format
First and only argument is an integer array denoting the preorder traversal of binary tree.

Output Format
Return a string "YES" if true else "NO".

Example Input
Input 1:
A : [4, 10, 5, 8]

Input 2:
A : [1, 5, 6, 4]

# Example Output

Output 1:
"YES"

Output 2:
"NO"

# Example Explanation

Explanation 1:
The possible BST is:

	4
	 \
	 10
	 /
	 5
	  \
	  8

Explanation 2:
There is no possible BST which have the above preorder traversal.
*/
func IsBstWithOneChild(A []int) string {
	if len(A) <= 2 {
		return "YES"
	}
	localMax := math.MinInt
	localMin := math.MaxInt
	for i := len(A) - 1; i >= 0; i-- {
		//idea is that for the current element, the right subarray should either be smaller or greater
		if i < len(A)-1 && A[i] < localMax && A[i] > localMin {
			return "NO"
		}
		localMax = max(localMax, A[i])
		localMin = min(localMin, A[i])
	}
	return "YES"
}
