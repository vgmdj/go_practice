package Combination_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1 1 1 5 7

		 1
     1        1      5    7
 1   5  7   5  7     7
5 7  7      7
*/
func TestCombinationSum(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]int{{7}}, combinationSum2([]int{2, 3, 6, 7}, 7))

	ast.Equal([][]int{{1, 3, 4}}, combinationSum2([]int{1, 2, 3, 4}, 8))

	ast.Equal([][]int{
		{1, 4},
		{2, 3},
	}, combinationSum2([]int{1, 2, 3, 4}, 5))

	ast.EqualValues([][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))

	ast.EqualValues([][]int{
		{1, 2, 2},
		{5},
	}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))

}
