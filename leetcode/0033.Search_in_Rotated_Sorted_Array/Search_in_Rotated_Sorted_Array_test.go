package Search_in_Rotated_Sorted_Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	ast.Equal(-1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))

}
