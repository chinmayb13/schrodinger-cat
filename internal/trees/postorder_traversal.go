package trees

/*
Problem Description
Given a binary tree, return the Postorder traversal of its nodes values.
NOTE: Using recursion is not allowed.

Problem Constraints
1 <= number of nodes <= 105

Input Format
First and only argument is root node of the binary tree, A.

Output Format
Return an integer array denoting the Postorder traversal of the given binary tree.

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
[3, 2, 1]

Output 2:
[6, 3, 2, 1]

Example Explanation
Explanation 1:
The Preoder Traversal of the given tree is [3, 2, 1].

Explanation 2:
The Preoder Traversal of the given tree is [6, 3, 2, 1].
*/
func PostorderTraversalAlt(A *treeNode) []int {
	cur := A
	var s stack
	var pre *treeNode
	var postOrderArr []int

	for (cur != nil) || (len(s) > 0) {
		//Push all left nodes into the stack till it hits NULL
		for cur != nil {
			s = s.push(cur)
			cur = cur.left
		}

		//get leftmost node and check if the right child is already processed or absent
		cur = s.top()
		//processed or absent
		if cur.right == nil || cur.right == pre {
			//consume the data
			postOrderArr = append(postOrderArr, cur.value)
			s = s.pop()
			//save the lastly processed node
			pre = cur
			cur = nil
			//right node present
		} else {
			cur = cur.right
		}

	}
	return postOrderArr
}
