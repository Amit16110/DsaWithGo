package stackandqueue

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

func (s *Stack) Push(val interface{}) {
	s.list.PushBack(val)
}
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	e := s.list.Back()
	s.list.Remove(e)
	return e.Value, true
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
