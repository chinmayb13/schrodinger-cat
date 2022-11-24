package trees

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

type stack []*treeNode

func (s stack) push(node *treeNode) stack {
	return append(s, node)

}
func (s stack) pop() stack {
	return s[:len(s)-1]
}

func (s stack) top() *treeNode {
	return s[len(s)-1]
}
