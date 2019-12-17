package _039_Combination_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))

	ast.Equal([][]int{
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 2},
		[]int{1, 1, 1, 1, 1, 3},
		[]int{1, 1, 1, 1, 2, 2},
		[]int{1, 1, 1, 1, 4},
		[]int{1, 1, 1, 2, 3},
		[]int{1, 1, 2, 2, 2},
		[]int{1, 1, 2, 4},
		[]int{1, 1, 3, 3},
		[]int{1, 2, 2, 3},
		[]int{1, 3, 4},
		[]int{2, 2, 2, 2},
		[]int{2, 2, 4},
		[]int{2, 3, 3},
		[]int{4, 4},
	},
		combinationSum([]int{1, 2, 3, 4}, 8))

}
