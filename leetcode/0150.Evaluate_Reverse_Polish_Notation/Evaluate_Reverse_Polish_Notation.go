package Evaluate_Reverse_Polish_Notation

import (
	"container/list"
	"strconv"
	"strings"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	s := newStack()
	for _, v := range tokens {
		if !strings.Contains("+-*/", v) {
			value, _ := strconv.Atoi(v)
			s.Push(value)
			continue
		}

		second := s.Pop()
		first := s.Pop()
		newNumber := getResult(first, second, v)
		s.Push(newNumber)
	}

	return s.Pop()
}

func getResult(first, second int, op string) int {
	switch op {
	case "+":
		return first + second

	case "-":
		return first - second

	case "*":
		return first * second

	case "/":
		return first / second

	default:
		return 0
	}
}

type stack struct {
	list *list.List
}

func newStack() *stack {
	return &stack{
		list.New(),
	}
}

func (s *stack) Pop() int {
	b := s.list.Back()
	res := b.Value.(int)
	s.list.Remove(b)

	return res
}

func (s *stack) Push(x int) {
	s.list.PushBack(x)
}

func (s *stack) Len() int {
	return s.list.Len()
}
