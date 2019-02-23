package Find_First_and_Last_Position_of_Element_in_Sorted_Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]int{0,2},searchRange([]int{3,3,3},3))
	ast.Equal([]int{5,12}, searchRange([]int{1, 2, 3, 4, 5, 8, 8, 8, 8, 8, 8, 8, 8, 9, 10, 11, 12, 13}, 8))
	ast.Equal([]int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 9}, 6))

}
