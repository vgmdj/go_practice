package Min_Stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ast := assert.New(t)

	ms.Push(6)
	ms.Push(1)
	ms.Push(-4)
	ast.Equal(-4, ms.GetMin())
	ms.Push(3)
	ms.Pop()
	ast.Equal(-4, ms.Top())
	ms.Pop()
	ast.Equal(1, ms.GetMin())

}
