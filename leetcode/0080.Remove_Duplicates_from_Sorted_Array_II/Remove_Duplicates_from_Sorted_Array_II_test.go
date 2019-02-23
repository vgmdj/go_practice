package Remove_Duplicates_from_Sorted_Array_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2,removeDuplicates([]int{1,2}))
	ast.Equal(5, removeDuplicates([]int{1, 1, 2, 2, 2, 2, 4}))
}
