package Search_in_Rotated_Sorted_Array_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(false, search([]int{}, 2))
	ast.Equal(true, search([]int{3, 1}, 1))
	ast.Equal(true, search([]int{2, 3, 2, 2, 2}, 3))
	ast.Equal(true, search([]int{2, 5, 6, 0, 0, 1, 2}, 1))
	ast.Equal(true, search([]int{4, 5, 5, 6, 6, 0, 1, 2, 3}, 2))

}
