package trees2

import "math"

/*
Given a binary tree, return a 2-D array with vertical order traversal of it. Go through the example and image for more details.

NOTE: If 2 Tree Nodes shares the same vertical level then the one with lesser depth will come first.

Problem Constraints
0 <= number of nodes <= 105

Input Format
First and only arument is a pointer to the root node of binary tree, A.

Output Format
Return a 2D array denoting the vertical order traversal of tree as shown.

Example Input
Input 1:

	     6
	   /   \
	  3     7
	 / \     \
	2   5     9

Input 2:

	     1
	   /   \
	  3     7
	 /       \
	2         9

Example Output
Output 1:

	[
	   [2],
	   [3],
	   [6, 5],
	   [7],
	   [9]
	]

Output 2:

	[
	   [2],
	   [3],
	   [1],
	   [7],
	   [9]
	]

Example Explanation
Explanation 1:
First row represent the verical line 1 and so on.
*/
func VerticalOrderTraversal(A *treeNode) [][]int {
	var ans [][]int
	var arr []int
	//map to store elements against a given vertical level
	levelMap := make(map[int][]*treeNode)
	q := newQueue()
	q.enqueueBack(newNodeInfo(A, 0))
	maxLevel, minLevel := math.MinInt, math.MaxInt
	for q.size() > 0 {
		ni := q.frontNodeInfo()
		q.dequeueFront()
		//add the element to the map against the corresponding level
		levelMap[ni.level] = append(levelMap[ni.level], ni.node)
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
		arr = nil
		for j := range levelMap[i] {
			arr = append(arr, levelMap[i][j].value)
		}
		ans = append(ans, arr)
	}
	return ans

}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}

type nodeInfo struct {
	node  *treeNode
	level int
}

func newNodeInfo(node *treeNode, level int) *nodeInfo {
	return &nodeInfo{
		node:  node,
		level: level,
	}
}
