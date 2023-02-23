package trees5

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
	next  *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}
