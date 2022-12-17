package stackandqueue

type (
	stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// create a new stack.
func NewStack() *stack {
	return &stack{nil, 0}
}

func (s *stack) len() int {
	return s.length
}

func (s *stack) peek() interface{} {
	if s.length == 0 {
		return nil
	}

	return s.top.value
}

func (s *stack) pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *stack) push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
