package Permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1 1 1 3
             1
  1  	 1       3
1  3   1  3	   1   1
3  1   3  1    1   1
*/

func TestPermutations(t *testing.T) {
	ast := assert.New(t)

	ast.EqualValues([][]int{
		{1},
	}, permuteUnique([]int{1}))

	ast.EqualValues([][]int{
		{1, 2},
		{2, 1},
	}, permuteUnique([]int{1, 2}))

	ast.EqualValues([][]int{
		{1, 1}}, permuteUnique([]int{1, 1}))

	ast.EqualValues([][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}, permuteUnique([]int{1, 2, 3}))

	ast.EqualValues([][]int{
		{1, 1, 1, 3},
		{1, 1, 3, 1},
		{1, 3, 1, 1},
		{3, 1, 1, 1},
	}, permuteUnique([]int{1, 1, 1, 3}))

	ast.EqualValues([][]int{
		[]int{1, 2, 3, 4}, []int{1, 2, 4, 3}, []int{1, 3, 2, 4}, []int{1, 3, 4, 2}, []int{1, 4, 2, 3}, []int{1, 4, 3, 2},
		[]int{2, 1, 3, 4}, []int{2, 1, 4, 3}, []int{2, 3, 1, 4}, []int{2, 3, 4, 1}, []int{2, 4, 1, 3}, []int{2, 4, 3, 1},
		[]int{3, 1, 2, 4}, []int{3, 1, 4, 2}, []int{3, 2, 1, 4}, []int{3, 2, 4, 1}, []int{3, 4, 1, 2}, []int{3, 4, 2, 1},
		[]int{4, 1, 2, 3}, []int{4, 1, 3, 2}, []int{4, 2, 1, 3}, []int{4, 2, 3, 1}, []int{4, 3, 1, 2}, []int{4, 3, 2, 1}},
		permuteUnique([]int{1, 2, 3, 4}))

}
