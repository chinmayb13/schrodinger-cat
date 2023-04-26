package tries2

/*
Problem Description
Given a binary tree A, flatten it to a linked list in-place.
The left child of all nodes should be NULL.

Problem Constraints
1 <= size of tree <= 100000

Input Format
First and only argument is the head of tree A.

Output Format
Return the linked-list after flattening.

Example Input
Input 1:

	  1
	 / \
	2   3

Input 2:

	    1
	   / \
	  2   5
	 / \   \
	3   4   6

Example Output
Output 1:
1

	\
	 2
	  \
	   3

Output 2:
1

	\
	 2
	  \
	   3
	    \
	     4
	      \
	       5
	        \
	         6

Example Explanation
Explanation 1:
Tree flattening looks like this.

Explanation 2:
Tree flattening looks like this.
*/
func Flatten(root *treeNode) *treeNode {
	newRoot, _ := flattenIntoLL(root)
	return newRoot
}

func flattenIntoLL(root *treeNode) (*treeNode, *treeNode) {
	if root == nil {
		return nil, nil
	}
	left, lTail := flattenIntoLL(root.left)
	right, rTail := flattenIntoLL(root.right)

	//as asked in the question, mark the left child to null
	root.left = nil

	//if left subtree is present, make the tail of left subtree point to the right head
	if lTail != nil {
		lTail.right = right
	}

	//if left subtree is present, make the root of right point to it, else point to right subtree
	if left != nil {
		root.right = left
	} else {
		root.right = right
	}

	//if right subtree is present, return root and right tail
	if rTail != nil {
		return root, rTail
		//else if left subtree is present, return root and left tail
	} else if lTail != nil {
		return root, lTail
	}
	//else return only one node
	return root, root
}
