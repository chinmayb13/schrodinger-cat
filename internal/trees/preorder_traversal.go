package trees

/*
Problem Description
Given a binary tree, return the preorder traversal of its nodes values.
NOTE: Using recursion is not allowed.

Problem Constraints
1 <= number of nodes <= 105

Input Format
First and only argument is root node of the binary tree, A.

Output Format
Return an integer array denoting the preorder traversal of the given binary tree.

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
[1, 2, 3]

Output 2:
[1, 6, 2, 3]

Example Explanation
Explanation 1:
The Preoder Traversal of the given tree is [1, 2, 3].

Explanation 2:
The Preoder Traversal of the given tree is [1, 6, 2, 3].
*/

func PreorderTraversal(A *treeNode) []int {
	cur := A
	var s stack
	var preOrderArr []int

	//check till all the nodes are processed
	for cur != nil || len(s) > 0 {
		for cur != nil {
			//since fashion is D->L->R, consume the data as soon as it is available
			preOrderArr = append(preOrderArr, cur.value)

			//keep pushing the intermediate nodes till we reach the leftmost node
			s = s.push(cur)
			cur = cur.left
		}
		//reached leftmost node
		cur = s.top()
		//pop the node
		s = s.pop()
		//process right node
		cur = cur.right

	}
	return preOrderArr
}
