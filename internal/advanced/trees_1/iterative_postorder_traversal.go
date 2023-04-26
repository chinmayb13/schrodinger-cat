package trees1

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
func PostorderTraversal(A *treeNode) []int {
	var ans []int
	st := newStack()
	temp := A
	for temp != nil || st.size() > 0 {
		if temp != nil {
			st.push(temp)
			temp = temp.left
		} else {
			temp = st.topNode()
			if temp.right == nil || (len(ans) > 0 && (ans[len(ans)-1] == temp.right.value)) {
				ans = append(ans, temp.value)
				st.pop()
				temp = nil
			} else {
				temp = temp.right
			}
		}
	}
	return ans
}
