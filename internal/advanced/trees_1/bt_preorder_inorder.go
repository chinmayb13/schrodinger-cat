package trees1

/*
Problem Description
Given preorder and inorder traversal of a tree, construct the binary tree.

NOTE: You may assume that duplicates do not exist in the tree.

Problem Constraints
1 <= number of nodes <= 105

Input Format
First argument is an integer array A denoting the preorder traversal of the tree.
Second argument is an integer array B denoting the inorder traversal of the tree.

Output Format
Return the root node of the binary tree.

Example Input
Input 1:
A = [1, 2, 3]
B = [2, 1, 3]

Input 2:
A = [1, 6, 2, 3]
B = [6, 1, 3, 2]

Example Output
Output 1:

	  1
	 / \
	2   3

Output 2:

	  1
	 / \
	6   2
	   /
	  3

Example Explanation
Explanation 1:
Create the binary tree and return the root node of the tree.
*/
func BuildBTree(A []int, B []int) *treeNode {
	size := len(A)
	inMap := make(map[int]int)
	for i := range B {
		inMap[B[i]] = i
	}
	return constructBTree(B, A, 0, size-1, 0, size-1, inMap)
}

func constructBTree(in, pre []int, inBeg, inEnd, preBeg, preEnd int, inMap map[int]int) *treeNode {
	if inBeg > inEnd {
		return nil
	}
	idx := inMap[pre[preBeg]]
	length := idx - inBeg
	root := treeNode_new(pre[preBeg])
	root.left = constructBTree(in, pre, inBeg, idx-1, preBeg+1, preBeg+length, inMap)
	root.right = constructBTree(in, pre, idx+1, inEnd, preBeg+length+1, preEnd, inMap)
	return root
}
