package trees2

import "math"

/*
Problem Description
Given a binary tree of integers denoted by root A. Return an array of integers representing the top view of the Binary tree.
The top view of a Binary Tree is a set of nodes visible when the tree is visited from the top.
Return the nodes in any order.

Problem Constraints
1 <= Number of nodes in binary tree <= 100000
0 <= node values <= 10^9

Input Format
First and only argument is head of the binary tree A.

Output Format
Return an array, representing the top view of the binary tree.

Example Input
Input 1:

	       1
	     /   \
	    2    3
	   / \  / \
	  4   5 6  7
	 /
	8

Input 2:

	  1
	 /  \
	2    3
	 \
	  4
	   \
	    5

Example Output
Output 1:
[1, 2, 4, 8, 3, 7]

Output 2:
[1, 2, 3]

Example Explanation
Explanation 1:
Top view is described.

Explanation 2:
Top view is described.
*/
func GetTopView(A *treeNode) []int {
	//map to store elements against a given vertical level
	nodeMap := make(map[int][]*treeNode)
	var ans []int
	q := newQueue()
	maxLevel, minLevel := math.MinInt, math.MaxInt
	q.enqueueBack(newNodeInfo(A, 0))
	for q.size() > 0 {
		ni := q.frontNodeInfo()
		q.dequeueFront()
		//add the element to the map against the corresponding level
		nodeMap[ni.level] = append(nodeMap[ni.level], ni.node)
		maxLevel = max(maxLevel, ni.level)
		minLevel = min(minLevel, ni.level)
		//add left child with one lesser level
		if ni.node.left != nil {
			q.enqueueBack(newNodeInfo(ni.node.left, ni.level-1))
		}
		//add right child with one greater level
		if ni.node.right != nil {
			q.enqueueBack(newNodeInfo(ni.node.right, ni.level+1))
		}
	}

	for i := minLevel; i <= maxLevel; i++ {
		//add only the first element of every vertical level
		ans = append(ans, nodeMap[i][0].value)
	}
	return ans
}
