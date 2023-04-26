package trees2

/*
Given a binary tree, return the values of its boundary in anti-clockwise direction starting from the root. Boundary includes left boundary, leaves, and right boundary in order without duplicate nodes.
Left boundary is defined as the path from the root to the left-most node. Right boundary is defined as the path from the root to the right-most node. If the root doesn't have left subtree or right subtree, then the root itself is left boundary or right boundary. Note this definition only applies to the input binary tree, and not applies to any subtrees.
The left-most node is defined as a leaf node you could reach when you always firstly travel to the left subtree if exists. If not, travel to the right subtree. Repeat until you reach a leaf node.
The right-most node is also defined by the same way with left and right exchanged.
Return an array of integers denoting the boundary values of tree in anti-clockwise order.

For Example
Input 1:

	    _____1_____
	   /           \
	  2             3
	 / \            /
	4   5          6
	   / \        / \
	  7   8      9  10

Output 1:

	[1, 2, 4, 7, 8, 9, 10, 6, 3]
	Explanation 1:
	    The left boundary are node 1,2,4. (4 is the left-most node according to definition)
	    The leaves are node 4,7,8,9,10.
	    The right boundary are node 1,3,6,10. (10 is the right-most node).
	    So order them in anti-clockwise without duplicate nodes we have [1,2,4,7,8,9,10,6,3].

Input 2:

	    1
	   / \
	  2   3
	 / \  / \
	4   5 6  7

Output 2:

	[1, 2, 4, 5, 6, 7, 3]
*/
func GetTreeBoundary(A *treeNode) []int {
	var leftView, rightView []int
	var leaves []int
	q := newQueue()
	q.enqueueBack(A)
	for q.size() > 0 {
		sz := q.size()
		for i := 0; i < sz; i++ {
			node := q.frontNode()
			q.dequeueFront()
			//store left boundary
			if i == 0 && !isLeaf(node) {
				leftView = append(leftView, node.value)
			}
			//store right boundary
			if sz > 1 && i == sz-1 && !isLeaf(node) {
				rightView = append(rightView, node.value)
			}
			if node.left != nil {
				q.enqueueBack(node.left)
			}
			if node.right != nil {
				q.enqueueBack(node.right)
			}
		}
	}
	//store lower boundary
	getLN(A, &leaves)
	leftView = append(leftView, leaves...)
	//reverse right boundary because of anticlockwise fashion
	for i := len(rightView) - 1; i >= 0; i-- {
		leftView = append(leftView, rightView[i])
	}
	return leftView
}

func getLN(A *treeNode, leaves *[]int) int {
	if A == nil {
		return -1
	}
	count := getLN(A.left, leaves) + getLN(A.right, leaves) + 2
	if count == 0 {
		*leaves = append(*leaves, A.value)
	}
	return count
}

func isLeaf(n *treeNode) bool {
	return n.left == nil && n.right == nil
}
