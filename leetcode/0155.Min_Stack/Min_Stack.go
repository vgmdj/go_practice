package Min_Stack

import "math"

type MinStack struct {
	values    []int
	minValues []int
	index     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		values:    make([]int, 0),
		minValues: []int{math.MaxInt32},
	}
}

func (s *MinStack) Push(x int) {
	s.values = append(s.values, x)
	if x <= s.minValues[len(s.minValues)-1] {
		s.minValues = append(s.minValues, x)
	}

}

func (s *MinStack) Pop() {
	if len(s.values) < 1 {
		return
	}

	if s.values[len(s.values)-1] == s.minValues[len(s.minValues)-1] {
		s.minValues = s.minValues[:len(s.minValues)-1]
	}
	s.values = s.values[:len(s.values)-1]

}

func (s *MinStack) Top() int {
	if len(s.values) < 1 {
		return -1
	}

	return s.values[len(s.values)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.minValues) < 2 {
		return -1
	}
	return s.minValues[len(s.minValues)-1]
}
