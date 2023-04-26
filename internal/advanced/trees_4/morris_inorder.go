package trees4

/*
Problem Description
Given a binary tree, return the inorder traversal of its nodes' values.

NOTE: Using recursion and stack are not allowed.

Problem Constraints
1 <= number of nodes <= 105

Input Format
First and only argument is root node of the binary tree, A.

Output Format
Return an integer array denoting the inorder traversal of the given binary tree.

Example Input
Input 1:

	1
	 \
	  2
	 /
	3

Input 2:

	  1
	 / \
	6   2
	   /
	  3

Example Output
Output 1:
[1, 3, 2]

Output 2:
[6, 1, 3, 2]

Example Explanation
Explanation 1:
The Inorder Traversal of the given tree is [1, 3, 2].

Explanation 2:
The Inorder Traversal of the given tree is [6, 1, 3, 2].
*/
func MorrisIOT(A *treeNode) []int {
	var ans []int
	cur := A
	var prev *treeNode
	for cur != nil {
		//pick the left subtree
		temp := cur.left
		prev = nil
		//go to most right of the left subtree
		for temp != nil && temp != cur {
			prev = temp
			temp = temp.right
		}
		//if the link is not created in leftsubtree, create the link and move the current to left
		if prev != nil && prev.right == nil {
			prev.right = cur
			cur = cur.left
			//if the link is already created, process the node, break the link and move the current to right
		} else if prev != nil && prev.right == cur {
			prev.right = nil
			ans = append(ans, cur.value)
			cur = cur.right
			//no left subtree is present
		} else {
			ans = append(ans, cur.value)
			cur = cur.right
		}
	}
	return ans
}
