package Kth_Largest_Element_in_a_Stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	ast := assert.New(t)

	k := 2
	arr := []int{0}
	kthLargest := Constructor(k, arr)

	ast.Equal(-1, kthLargest.Add(-1))
	ast.Equal(0, kthLargest.Add(1))
	ast.Equal(0, kthLargest.Add(-2))
	ast.Equal(0, kthLargest.Add(-4))
	ast.Equal(1, kthLargest.Add(3))

	k = 3
	arr = []int{5, -1}
	kthLargest = Constructor(k, arr)

	ast.Equal(-1, kthLargest.Add(2))
	ast.Equal(1, kthLargest.Add(1))
	ast.Equal(1, kthLargest.Add(-1))
	ast.Equal(2, kthLargest.Add(3))
	ast.Equal(3, kthLargest.Add(4))

	k = 3
	arr = []int{4, 5, 8, 2}
	kthLargest = Constructor(k, arr)

	ast.Equal(4, kthLargest.Add(3))
	ast.Equal(5, kthLargest.Add(5))
	ast.Equal(5, kthLargest.Add(10))
	ast.Equal(8, kthLargest.Add(9))
	ast.Equal(8, kthLargest.Add(4))

	k = 1
	arr = []int{2}
	kthLargest = Constructor(k, arr)

	ast.Equal(2, kthLargest.Add(2))
	ast.Equal(3, kthLargest.Add(3))
	ast.Equal(7, kthLargest.Add(7))
	ast.Equal(7, kthLargest.Add(1))

}
