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

	for i := range right {
		ast.Equal(nums[i], right[i])
	}

}

func TestSelect(t *testing.T) {
	ast := assert.New(t)

	nums := []int{8, 6, 4, 3, 2, 2, 1, 9}
	selectSort(nums)

	right := []int{1, 2, 2, 3, 4, 6, 8, 9}

	for i := range right {
		ast.Equal(nums[i], right[i])
	}

}

func TestInsert(t *testing.T) {
	ast := assert.New(t)

	nums := []int{8, 6, 4, 3, 2, 2, 1, 9}
	insertSort(nums)

	right := []int{1, 2, 2, 3, 4, 6, 8, 9}

	for i := range right {
		ast.Equal(nums[i], right[i])
	}

}

func TestHeap(t *testing.T) {
	ast := assert.New(t)

	nums := []int{8, 6, 4, 3, 2, 2, 1, 9}
	HeapSort(nums)

	right := []int{1, 2, 2, 3, 4, 6, 8, 9}

	for i := range right {
		ast.Equal(nums[i], right[i])
	}

}
