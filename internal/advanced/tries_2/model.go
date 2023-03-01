package tries2

type btNode struct {
	value int
	left  *btNode
	right *btNode
}

func newbtNode() *btNode {
	return &btNode{}
}

type info struct {
	xor int
	beg int
}

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func newTreeNode(val int) *treeNode {
	return &treeNode{
		value: val,
	}
}
