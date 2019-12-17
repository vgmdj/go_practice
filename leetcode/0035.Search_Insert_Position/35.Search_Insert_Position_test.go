package Search_Insert_Position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(searchInsert([]int{1, 3, 5, 6}, 5), 2)
	ast.Equal(searchInsert([]int{1, 3, 5, 6}, 2), 1)
	ast.Equal(searchInsert([]int{1, 3, 5, 6}, 7), 4)
	ast.Equal(searchInsert([]int{1, 3, 5, 6}, 0), 0)

}

func TestSearchInsertII(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(searchInsert2([]int{1, 3, 5, 6}, 5), 2)
	ast.Equal(searchInsert2([]int{1, 3, 5, 6}, 2), 1)
	ast.Equal(searchInsert2([]int{1, 3, 5, 6}, 7), 4)
	ast.Equal(searchInsert2([]int{1, 3, 5, 6}, 0), 0)

}
