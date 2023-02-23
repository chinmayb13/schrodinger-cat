package trees5

/*
Problem Description
Given a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Problem Constraints
1 <= number of nodes <= 105

Input Format
First and only argument is the root node of the binary tree.

Output Format
Return 0 / 1 ( 0 for false, 1 for true ).

Example Input
Input 1:
    1
   / \
  2   2
 / \ / \
3  4 4  3

Input 2:
    1
   / \
  2   2
   \   \
   3    3

Example Output

Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
The above binary tree is symmetric.

Explanation 2:
The above binary tree is not symmetric.
*/
func IsSymmetric(A *treeNode) int {
	ok := isSymmetric(A.left, A.right)
	if ok {
		return 1
	}
	return 0
}

func isSymmetric(root1, root2 *treeNode) bool {
	if root1 == nil && root2 == nil {
		return true
		//comparison should be done between left most and right most, second left and second right
	} else if root1 != nil && root2 != nil && root1.value == root2.value {
		return isSymmetric(root1.left, root2.right) && isSymmetric(root1.right, root2.left)
	}
	return false
}
