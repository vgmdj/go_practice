package _4Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test4Sum(t *testing.T) {
	ast := assert.New(t)

	ast.EqualValues([][]int{
		{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		fourSum([]int{1, 0, -1, 0, -2, 2}, 0))

	ast.EqualValues([][]int{
		[]int{-4, 1, 4, 4}, []int{-2, 0, 3, 4}, []int{0, 0, 1, 4}, []int{0, 1, 1, 3}},
		fourSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}, 5))

}
