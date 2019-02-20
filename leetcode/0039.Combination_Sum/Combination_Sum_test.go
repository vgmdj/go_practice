package _039_Combination_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]int{{3, 2, 2}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))

	ast.Equal([][]int{
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		[]int{2, 2, 2, 2},
		[]int{2, 2, 2, 1, 1},
		[]int{2, 2, 1, 1, 1, 1},
		[]int{2, 1, 1, 1, 1, 1, 1},
		[]int{3, 3, 2},
		[]int{3, 3, 1, 1},
		[]int{3, 2, 2, 1},
		[]int{3, 2, 1, 1, 1},
		[]int{3, 1, 1, 1, 1, 1},
		[]int{4, 4},
		[]int{4, 3, 1},
		[]int{4, 2, 2},
		[]int{4, 2, 1, 1},
		[]int{4, 1, 1, 1, 1}},
		combinationSum([]int{1, 2, 3, 4}, 8))

}
