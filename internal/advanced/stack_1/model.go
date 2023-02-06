package stack1

import (
	"container/list"
	"strconv"
)

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

func (s *stack) topByte() byte {
	top := s.top()
	if top != nil {
		return top.(byte)
	}
	return 0
}

func (s *stack) topInt() int {
	top := s.top()
	if top != nil {
		return top.(int)
	}
	return 0
}

func (s *stack) topStrInt() int {
	top := s.top()
	if top != nil {
		val, _ := strconv.Atoi(top.(string))
		return val
	}
	return 0
}

func (s *stack) topBool() bool {
	top := s.top()
	if top != nil {
		return top.(bool)
	}
	return false
}
