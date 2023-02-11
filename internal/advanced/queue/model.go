package queue

import (
	"container/list"
)

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

func (s *queue) backByte() byte {
	back := s.back()
	if back != nil {
		return back.(byte)
	}
	return 0
}

func (s *queue) frontByte() byte {
	front := s.front()
	if front != nil {
		return front.(byte)
	}
	return 0
}

func (s *queue) backInt() int {
	back := s.back()
	if back != nil {
		return back.(int)
	}
	return 0
}

func (s *queue) frontInt() int {
	front := s.front()
	if front != nil {
		return front.(int)
	}
	return 0
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

func (s *stack) topInt() int {
	top := s.top()
	if top != nil {
		return top.(int)
	}
	return 0
}
