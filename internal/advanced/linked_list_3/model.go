package linkedlist3

type listNode struct {
	value int
	next  *listNode
}

func listNode_new(val int) *listNode {
	var node *listNode = new(listNode)
	node.value = val
	node.next = nil
	return node
}

type listNodeAlt struct {
	val   int
	right *listNodeAlt
	down  *listNodeAlt
}

func listNodeAlt_new(val int) *listNodeAlt {
	var node *listNodeAlt = new(listNodeAlt)
	node.val = val
	node.right = nil
	node.down = nil
	return node
}
