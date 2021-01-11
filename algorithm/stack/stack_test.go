package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	ast := assert.New(t)

	s := NewStack()

	s.Push(1)
	ast.Equal(1, s.Len())

	v1, ok := s.Pop()
	ast.Equal(true, ok)
	ast.Equal(1, v1)

	s.Push(2)
	s.Push(3)

	v3, ok := s.Pop()
	ast.Equal(true, ok)
	ast.Equal(3, v3)

	v2, ok := s.Pop()
	ast.Equal(true, ok)
	ast.Equal(2, v2)

	n, ok := s.Pop()
	ast.Equal(false, ok)
	ast.Equal(-1, n)
}
