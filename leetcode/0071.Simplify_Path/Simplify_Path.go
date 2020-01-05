package Simplify_Path

import "strings"

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	s := NewStack()

	for _, v := range dirs {
		switch v {
		case "", ".":

		case "..":
			s.Pop()

		default:
			s.Push(v)

		}
	}

	return getResult(s)

}

func getResult(s *stack) string {
	return "/" + strings.Join(s.data[:s.index+1], "/")
}

type stack struct {
	data  []string
	index int
}

func NewStack() *stack {
	return &stack{
		data:  make([]string, 0),
		index: -1,
	}
}

func (s *stack) Push(str string) {
	s.index++
	if len(s.data) <= s.index{
		s.data = append(s.data,str)
		return
	}

	s.data[s.index] = str
}

func (s *stack) Pop() string {
	if s.index < 0 {
		return ""
	}

	s.index--

	return s.data[s.index+1]
}

func (s *stack) Len() int {
	return s.index + 1
}
