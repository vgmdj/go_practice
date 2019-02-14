package _3Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3Sum(t *testing.T) {
	ast := assert.New(t)

	ast.EqualValues([][]int{
		{-1, -1, 2}, {-1, 0, 1}},
		threeSum([]int{-1, 0, 1, 2, -1, -4}))

	ast.EqualValues([][]int{
		{-5, 1, 4}, {-5, 2, 3}, {-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3},
		{-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		threeSum([]int{1, 2, 3, 4, 0, -3, -2, -1, -1, -1, -1, -4, -5}))

	ast.EqualValues([][]int{
		{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}},
		threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))

}
