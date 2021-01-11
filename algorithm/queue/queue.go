package queue

// Queue the structure of custom queue
type Queue struct {
	list []interface{}
}

// NewQueue return the object of queue
func NewQueue() *Queue {
	return &Queue{
		list: make([]interface{}, 0),
	}
}

// Push push the value to queue top
func (s *Queue) Push(v interface{}) {
	s.list = append(s.list, v)
}

// Pop pop the value of queue top
func (s *Queue) Pop() (interface{}, bool) {
	if len(s.list) == 0 {
		return -1, false
	}

	v := s.list[0]
	s.list = s.list[1:]

	return v, true
}

// Head return the head of queue and not pop
func (s *Queue) Head() (interface{}, bool) {
	if len(s.list) == 0 {
		return -1, false
	}

	return s.list[0], true
}

// Tail return the tail of queue and not pop
func (s *Queue) Tail() (interface{}, bool) {
	if len(s.list) == 0 {
		return -1, false
	}

	return s.list[len(s.list)-1], true

}

// Len return the length of queue
func (s *Queue) Len() int {
	return len(s.list)
}
