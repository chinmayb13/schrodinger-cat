package trees1

/*
Problem Description
Given the inorder and postorder traversal of a tree, construct the binary tree.
NOTE: You may assume that duplicates do not exist in the tree.

Problem Constraints
1 <= number of nodes <= 105

Input Format
First argument is an integer array A denoting the inorder traversal of the tree.
Second argument is an integer array B denoting the postorder traversal of the tree.

Output Format
Return the root node of the binary tree.

Example Input
Input 1:
A = [2, 1, 3]
B = [2, 3, 1]

Input 2:
A = [6, 1, 3, 2]
B = [6, 3, 2, 1]

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
func BuildTree(A []int, B []int) *treeNode {
	size := len(A)
	inMap := make(map[int]int)
	for i := range A {
		inMap[A[i]] = i
	}
	return constructTree(A, B, 0, size-1, 0, size-1, inMap)
}

func constructTree(in, post []int, inBeg, inEnd, postBeg, postEnd int, inMap map[int]int) *treeNode {
	if inBeg > inEnd {
		return nil
	}
	idx := inMap[post[postEnd]]
	length := idx - inBeg
	root := treeNode_new(post[postEnd])
	root.left = constructTree(in, post, inBeg, idx-1, postBeg, postBeg+length-1, inMap)
	root.right = constructTree(in, post, idx+1, inEnd, postBeg+length, postEnd-1, inMap)
	return root
}
