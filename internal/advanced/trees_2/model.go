package trees2

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

func (s *queue) enqueueFront(inp interface{}) {
	s.PushFront(inp)
}

func (s *queue) dequeueBack() {
	if s.Len() > 0 {
		s.Remove(s.Back())
	}
}

func (s *queue) dequeueFront() {
	if s.Len() > 0 {
		s.Remove(s.Front())
	}
}

func (s *queue) back() interface{} {
	if s.Len() > 0 {
		return s.Back().Value
	}
	return nil
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

func (s *queue) backNode() *treeNode {
	back := s.back()
	if back != nil {
		return back.(*treeNode)
	}
	return nil
}

func (s *queue) frontNode() *treeNode {
	front := s.front()
	if front != nil {
		return front.(*treeNode)
	}
	return nil
}

func (s *queue) frontNodeInfo() *nodeInfo {
	front := s.front()
	if front != nil {
		return front.(*nodeInfo)
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
