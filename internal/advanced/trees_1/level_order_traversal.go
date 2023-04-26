package trees1

/*
Problem Description
Given a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

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
   [9, 20],
   [15, 7]
 ]

Output 2:

 [
   [1]
   [6, 2]
   [3]
 ]

Example Explanation
Explanation 1:
Return the 2D array. Each row denotes the traversal of each level.
*/

// idea: maintain a queue to to retain the level wise traversal
func GetLevelOrder(A *treeNode) [][]int {
	var ans [][]int
	var tempAns []int
	row := 1
	nodeQueue := newQueue()
	levelQueue := newQueue()
	levelQueue.enqueueBack(1)
	nodeQueue.enqueueBack(A)
	for nodeQueue.size() > 0 {
		temp := nodeQueue.frontNode()
		level := levelQueue.frontInt()
		if level == row {
			tempAns = append(tempAns, temp.value)
		} else {
			ans = append(ans, tempAns)
			tempAns = nil
			row++
			tempAns = append(tempAns, temp.value)
		}
		nodeQueue.dequeueFront()
		levelQueue.dequeueFront()
		if temp.left != nil {
			nodeQueue.enqueueBack(temp.left)
			levelQueue.enqueueBack(level + 1)
		}
		if temp.right != nil {
			nodeQueue.enqueueBack(temp.right)
			levelQueue.enqueueBack(level + 1)
		}
	}
	ans = append(ans, tempAns)
	return ans

}
