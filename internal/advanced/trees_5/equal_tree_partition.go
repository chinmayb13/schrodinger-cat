package trees5

/*
Problem Description
Given a binary tree A. Check whether it is possible to partition the tree to two trees which have equal sum of values after removing exactly one edge on the original tree.

Problem Constraints
1 <= size of tree <= 100000
0 <= value of node <= 109

Input Format
First and only argument is head of tree A.

Output Format
Return 1 if the tree can be partitioned into two trees of equal sum else return 0.

Example Input
Input 1:

	    5
	   /  \
	  3    7
	 / \  / \
	4  6  5  6

Input 2:

	  1
	 / \
	2   10
	    / \
	   20  2

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
Remove edge between 5(root node) and 7:

	Tree 1 =                                               Tree 2 =
	                5                                                     7
	               /                                                     / \
	              3                                                     5   6
	             / \
	            4   6
	Sum of Tree 1 = Sum of Tree 2 = 18

Explanation 2:
The given Tree cannot be partitioned.
*/
func CanTreeBePartitioned(A *treeNode) int {
	sum := getSum(A)
	_, ok := checkSum(A, sum>>1)
	if ok {
		return 1
	}
	return 0
}

func getSum(root *treeNode) int {
	if root == nil {
		return 0
	}
	return getSum(root.left) + getSum(root.right) + root.value
}

func checkSum(root *treeNode, target int) (int, bool) {
	if root == nil {
		return 0, false
	}
	leftSum, leftOk := checkSum(root.left, target)
	if leftOk {
		return leftSum, true
	}
	rightSum, rightOk := checkSum(root.right, target)
	sum := root.value + leftSum + rightSum
	if rightOk || sum == target {
		return sum, true
	}
	return sum, false
}
