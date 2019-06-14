package Kth_Largest_Element_in_an_Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(5, findKthLargest3([]int{3,2,1,5,6,4}, 2))
	ast.Equal(8, findKthLargest3([]int{1, 2, 4, 6, 7, 8, 534, 4, 7, 63, 1}, 3))
}
