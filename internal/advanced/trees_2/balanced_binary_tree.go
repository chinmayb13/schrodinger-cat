package trees2

/*
Problem Description
Given a root of binary tree A, determine if it is height-balanced.

A height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Problem Constraints
1 <= size of tree <= 100000

Input Format
First and only argument is the root of the tree A.

Output Format
Return 0 / 1 ( 0 for false, 1 for true ) for this problem.

Example Input
Input 1:
    1
   / \
  2   3

Input 2:
       1
      /
     2
    /
   3

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
It is a complete binary tree.

Explanation 2:
Because for the root node, left subtree has depth 2 and right subtree has depth 0.
Difference = 2 > 1.
*/
func IsBalanced(A *treeNode) int {
	_, ok := checkifBalanced(A)
	if ok {
		return 1
	}
	return 0
}

func checkifBalanced(root *treeNode) (int, bool) {
	if root == nil {
		return -1, true
	}
	lHeight, ok := checkifBalanced(root.left)
	if !ok {
		return -1, false
	}

	rHeight, ok1 := checkifBalanced(root.right)
	if !ok1 || abs(rHeight-lHeight) > 1 {
		return -1, false
	}
	return max(lHeight, rHeight) + 1, true
}

func abs(inp int) int {
	if inp < 0 {
		return -1 * inp
	}
	return inp
}
