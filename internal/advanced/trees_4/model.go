package trees4

import "container/list"

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

type stack struct {
	*list.List
}

func newStack() *stack {
	return &stack{
		List: list.New(),
	}
}

func (s *stack) push(inp interface{}) {
	s.PushBack(inp)
}

func (s *stack) pop() {
	if s.Len() > 0 {
		s.Remove(s.Back())
	}
}

func (s *stack) top() interface{} {
	if s.Len() > 0 {
		return s.Back().Value

	}
	return nil
}

func (s *stack) size() int {
	return s.Len()
}

func (s *stack) topNode() *treeNode {
	top := s.top()
	if top != nil {
		return top.(*treeNode)
	}
	return nil
}
