package trees5

/*
Problem Description
Given two binary trees, check if they are equal or not.
Two binary trees are considered equal if they are structurally identical and the nodes have the same value.

Problem Constraints
1 <= number of nodes <= 105

Input Format
The first argument is a root node of the first tree, A.
The second argument is a root node of the second tree, B.

Output Format
Return 0 / 1 ( 0 for false, 1 for true ) for this problem.

Example Input
Input 1:
   1       1
  / \     / \
 2   3   2   3

Input 2:
   1       1
  / \     / \
 2   3   3   3

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
Both trees are structurally identical and the nodes have the same value.

Explanation 2:
Values of the left child of the root node of the trees are different.
*/
func IsSameTree(A *treeNode, B *treeNode) int {
	ok := isSame(A, B)
	if ok {
		return 1
	}
	return 0
}

func isSame(root1, root2 *treeNode) bool {
	if root1 == nil && root2 == nil {
		return true
		//comparison should be done between left most and second left, second right and right most
	} else if root1 != nil && root2 != nil && root1.value == root2.value {
		return isSame(root1.left, root2.left) && isSame(root1.right, root2.right)
	}
	return false
}
