package trees1

import "container/list"

type queue struct {
	*list.List
}

func newQueue() *queue {
	return &queue{
		List: list.New(),
	}
}

func (s *queue) enqueueBack(inp interface{}) {
	s.PushBack(inp)
}

func (s *queue) dequeueFront() {
	if s.Len() > 0 {
		s.Remove(s.Front())
	}
}

func (s *queue) front() interface{} {
	if s.Len() > 0 {
		return s.Front().Value

	}
	return nil
}

func (s *queue) size() int {
	return s.Len()
}

func (s *queue) frontInt() int {
	front := s.front()
	if front != nil {
		return front.(int)
	}
	return 0
}

func (s *queue) frontNode() *treeNode {
	front := s.front()
	if front != nil {
		return front.(*treeNode)
	}
	return nil
}

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
