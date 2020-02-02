package Largest_Rectangle_in_Histogram

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	max := -1
	s := NewStack(len(heights))

	for i := 0; i < len(heights); i++ {
		// if heights[i] <= heights[s.Back()], s.Pop(), and calculate area
		for s.Len() > 0 && heights[i] <= heights[s.Back()] {
			p := s.Pop()
			area := heights[p] * (i - s.Back() - 1)
			if area > max {
				max = area
			}
		}

		s.Push(i)

	}

	right := len(heights)
	for s.Len() != 0 {
		left := s.Pop()
		area := heights[left] * (right - s.Back() - 1)
		if area > max {
			max = area
		}
	}

	return max

}

type stack struct {
	data  []int
	index int
}

func NewStack(cap int) *stack {
	return &stack{
		data:  make([]int, cap),
		index: -1,
	}
}

func (s *stack) Push(v int) {
	s.index++
	s.data[s.index] = v
}

func (s *stack) Pop() int {
	if s.index < 0 {
		return -1
	}

	v := s.data[s.index]
	s.index--

	return v
}

func (s *stack) Len() int {
	return s.index + 1
}

func (s *stack) Back() int {
	if s.index < 0 {
		return -1
	}

	return s.data[s.index]
}
