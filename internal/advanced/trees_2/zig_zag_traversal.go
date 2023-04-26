package trees2

/*
Problem Description
Given a binary tree, return the zigzag level order traversal of its nodes values. (ie, from left to right, then right to left for the next level and alternate between).

Problem Constraints
1 <= number of nodes <= 105

Input Format
First and only argument is root node of the binary tree, A.

Output Format
Return a 2D integer array denoting the zigzag level order traversal of the given binary tree.

Example Input
Input 1:

	  3
	 / \
	9  20
	  /  \
	 15   7

Input 2:

	  1
	 / \
	6   2
	   /
	  3

Example Output
Output 1:
[

	[3],
	[20, 9],
	[15, 7]

]

Output 2:
[

	[1]
	[2, 6]
	[3]

]

Example Explanation
Explanation 1:
Return the 2D array. Each row denotes the zigzag traversal of each level.
*/
func ZigzagLevelOrder(A *treeNode) [][]int {
	var arr []int
	var ans [][]int
	deque := newQueue()
	deque.enqueueBack(A)
	flag := 1
	for deque.size() > 0 {
		arr = nil
		sz := deque.size()
		//odd level traversal
		if flag == 1 {
			for i := 0; i < sz; i++ {
				node := deque.frontNode()
				deque.dequeueFront()
				arr = append(arr, node.value)
				if node.left != nil {
					deque.enqueueBack(node.left)
				}
				if node.right != nil {
					deque.enqueueBack(node.right)
				}
			}
			//even level traversal
		} else {
			for i := 0; i < sz; i++ {
				node := deque.backNode()
				deque.dequeueBack()
				arr = append(arr, node.value)
				if node.right != nil {
					deque.enqueueFront(node.right)
				}
				if node.left != nil {
					deque.enqueueFront(node.left)
				}
			}
		}
		ans = append(ans, arr)
		flag = flag ^ 1
	}
	return ans
}
