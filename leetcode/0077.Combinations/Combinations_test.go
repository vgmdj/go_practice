package Combinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]int{{1}, {2}, {3}}, combine(3, 1))

	ast.Equal([][]int{
		{1, 2},
	}, combine(2, 2))

	ast.Equal([][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}, combine(4, 2))

	ast.Equal([][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 4},
		[]int{1, 2, 5},
		[]int{1, 3, 4},
		[]int{1, 3, 5},
		[]int{1, 4, 5},
		[]int{2, 3, 4},
		[]int{2, 3, 5},
		[]int{2, 4, 5},
		[]int{3, 4, 5}},
		combine(5, 3))

}
