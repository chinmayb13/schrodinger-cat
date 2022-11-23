package trees

/*
Problem Description
Given a binary tree, return the inorder traversal of its nodes' values.
NOTE: Using recursion is not allowed.

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
func InorderTraversalAlt(A *treeNode) []int {
	cur := A
	var s stack
	var inOrderArr []int

	//check till all the elements are processed
	for cur != nil || len(s) > 0 {

		//go to most left of the tree and add intermediate elements in the stack
		for cur != nil {
			s = s.push(cur)
			cur = cur.left
		}
		//get the top element as soon as the cur is nil (meaning we are on leftmost node)
		cur = s.top()
		//store its value (meaning L->D->R)
		inOrderArr = append(inOrderArr, cur.value)
		//pop the node from stack
		s = s.pop()
		//process right node
		cur = cur.right

	}
	return inOrderArr
}

func InorderTraversalInRecursion(A *treeNode) []int {
	if A == nil {
		return nil
	}
	var inOrderArr []int
	inOrderArr = append(inOrderArr, InorderTraversalInRecursion(A.left)...)
	inOrderArr = append(inOrderArr, A.value)
	inOrderArr = append(inOrderArr, InorderTraversalInRecursion(A.right)...)
	return inOrderArr
}
