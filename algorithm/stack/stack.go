package stack

// Stack the structure of custom stack
type Stack struct {
	list []interface{}
}

// NewStack return the object of stack
func NewStack() *Stack {
	return &Stack{
		list: make([]interface{}, 0),
	}
}

// Push push the value to stack top
func (s *Stack) Push(v interface{}) {
	s.list = append(s.list, v)
}

// Pop pop the value of stack top
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.list) == 0 {
		return -1, false
	}

	v := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return v, true
}

// Top return the stack top value and not pop
func (s *Stack) Top() (interface{}, bool) {
	if len(s.list) == 0 {
		return -1, false
	}

	return s.list[len(s.list)-1], true
}

// Bottom return the bottom of stack and not pop
func (s *Stack) Bottom() (interface{}, bool) {
	if len(s.list) == 0 {
		return -1, false
	}

	return s.list[0], true
}

// Len return the length of stack
func (s *Stack) Len() int {
	return len(s.list)
}
