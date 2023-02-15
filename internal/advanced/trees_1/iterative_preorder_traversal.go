package trees1

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
	var ans []int
	st := newStack()
	temp := A
	for temp != nil || st.size() > 0 {
		if temp != nil {
			ans = append(ans, temp.value)
			st.push(temp)
			temp = temp.left
		} else {
			temp = st.topNode()
			st.pop()
			temp = temp.right
		}

	}
	return ans
}
