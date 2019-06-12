package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {
	ast := assert.New(t)

	nums := []int{8, 6, 4, 3, 2, 2, 1, 9}
	bubbleSort(nums)

	right := []int{1, 2, 2, 3, 4, 6, 8, 9}

	ast.Equal(right, nums)

}

func TestSelect(t *testing.T) {
	ast := assert.New(t)

	nums := []int{8, 6, 4, 3, 2, 2, 1, 9}
	selectSort(nums)

	right := []int{1, 2, 2, 3, 4, 6, 8, 9}

	ast.Equal(right, nums)

}

func TestInsert(t *testing.T) {
	ast := assert.New(t)

	nums := []int{8, 6, 4, 3, 2, 2, 1, 9}
	insertSort(nums)

	right := []int{1, 2, 2, 3, 4, 6, 8, 9}

	ast.Equal(right, nums)

}

func TestQuick(t *testing.T) {
	ast := assert.New(t)

	nums := []int{8, 6, 4, 3, 2, 2, 1, 9}
	quickSort(nums)

	right := []int{1, 2, 2, 3, 4, 6, 8, 9}

	ast.Equal(right, nums)

}

func TestHeap(t *testing.T) {
	ast := assert.New(t)

	nums := []int{8, 6, 4, 3, 2, 2, 1, 9}
	HeapSort(nums)

	right := []int{1, 2, 2, 3, 4, 6, 8, 9}

	ast.Equal(right, nums)

}
