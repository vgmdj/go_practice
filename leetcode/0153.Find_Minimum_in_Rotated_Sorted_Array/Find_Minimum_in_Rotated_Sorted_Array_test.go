package Find_Minimum_in_Rotated_Sorted_Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1, findMin([]int{3, 4, 5, 1, 2}))
	ast.Equal(0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))

}
