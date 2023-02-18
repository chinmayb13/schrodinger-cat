package trees2

/*
Problem Description
Given a binary tree of integers. Find the difference between the sum of nodes at odd level and sum of nodes at even level.
NOTE: Consider the level of root node as 1.

Problem Constraints
1 <= Number of nodes in binary tree <= 100000
0 <= node values <= 109

Input Format
First and only argument is a root node of the binary tree, A

Output Format
Return an integer denoting the difference between the sum of nodes at odd level and sum of nodes at even level.

Example Input
Input 1:
        1
      /   \
     2     3
    / \   / \
   4   5 6   7
  /
 8

Input 2:
        1
       / \
      2   10
       \
        4

Example Output
Output 1:
10

Output 2:
-7

Example Explanation
Explanation 1:
Sum of nodes at odd level = 23
Sum of ndoes at even level = 13

Explanation 2:
Sum of nodes at odd level = 5
Sum of ndoes at even level = 12
*/
func GetOddEvenLevelDiff(A *treeNode) int {
	var oddSum, evenSum int
	q := newQueue()
	q.enqueueBack(A)
	flag := 1
	for q.size() > 0 {
		sz := q.size()
		sum := 0
		for i := 0; i < sz; i++ {
			node := q.frontNode()
			q.dequeueFront()
			sum += node.value
			if node.left != nil {
				q.enqueueBack(node.left)
			}
			if node.right != nil {
				q.enqueueBack(node.right)
			}
		}
		if flag&1 > 0 {
			oddSum += sum
		} else {
			evenSum += sum
		}
		flag = flag ^ 1
	}
	return oddSum - evenSum
}
