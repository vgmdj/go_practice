package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	ast := assert.New(t)

	h := NewHeap(NewMaxHeap())
	test := []int{9, 4, 5, 6, 3, 2, 1, 8}
	for _, v := range test {
		h.Push(v)
	}

	ans := []int{9, 8, 6, 5, 4, 3, 2, 1}
	for _, c := range ans {
		v, ok := h.Pop()
		if !ok {
			t.Error("excepted true,but got false")
			return
		}

		ast.Equal(c, v)
	}

}

func TestMinHeap(t *testing.T) {
	ast := assert.New(t)

	h := NewHeap(NewMinHeap())
	test := []int{9, 4, 5, 6, 3, 2, 1, 8}
	for _, v := range test {
		h.Push(v)
	}

	ans := []int{1, 2, 3, 4, 5, 6, 8, 9}
	for _, c := range ans {
		v, ok := h.Pop()
		if !ok {
			t.Error("excepted true,but got false")
			return
		}

		ast.Equal(c, v)
	}

}
