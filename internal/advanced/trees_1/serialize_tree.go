package trees1

/*
Problem Description
Given the root node of a Binary Tree denoted by A. You have to Serialize the given Binary Tree in the described format.

Serialize means encode it into a integer array denoting the Level Order Traversal of the given Binary Tree.

NOTE:
In the array, the NULL/None child is denoted by -1.
For more clarification check the Example Input.

Problem Constraints
1 <= number of nodes <= 105

Input Format
Only argument is a A denoting the root node of a Binary Tree.

Output Format
Return an integer array denoting the Level Order Traversal of the given Binary Tree.

Example Input
Input 1:
           1
         /   \
        2     3
       / \
      4   5

Input 2:

            1
          /   \
         2     3
        / \     \
       4   5     6

Example Output
Output 1:
[1, 2, 3, 4, 5, -1, -1, -1, -1, -1, -1]

Output 2:
[1, 2, 3, 4, 5, -1, 6, -1, -1, -1, -1, -1, -1]

Example Explanation
Explanation 1:
The Level Order Traversal of the given tree will be [1, 2, 3, 4, 5 , -1, -1, -1, -1, -1, -1].
Since 3, 4 and 5 each has both NULL child we had represented that using -1.

Explanation 2:
The Level Order Traversal of the given tree will be [1, 2, 3, 4, 5, -1, 6, -1, -1, -1, -1, -1, -1].
Since 3 has left child as NULL while 4 and 5 each has both NULL child.
*/
func SerializeTree(A *treeNode) []int {
	var ans []int
	q := newQueue()
	q.enqueueBack(A)
	for q.size() > 0 {
		node := q.frontNode()
		ans = append(ans, node.value)
		if node.value != -1 {
			if node.left != nil {
				q.enqueueBack(node.left)
			} else {
				q.enqueueBack(treeNode_new(-1))
			}

			if node.right != nil {
				q.enqueueBack(node.right)
			} else {
				q.enqueueBack(treeNode_new(-1))
			}
		}
		q.dequeueFront()
	}
	return ans
}
