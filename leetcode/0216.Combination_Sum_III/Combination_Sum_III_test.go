package Combination_Sum_III

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]int{{9}}, combinationSum3(1, 9))

	ast.Equal([][]int{{1, 2, 4}},
		combinationSum3(3, 7))

	ast.Equal([][]int{
		{1, 2, 6},
		{1, 3, 5},
		{2, 3, 4},
	}, combinationSum3(3, 9))
}
