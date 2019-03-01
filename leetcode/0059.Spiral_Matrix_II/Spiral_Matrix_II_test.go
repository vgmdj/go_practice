package Spiral_Matrix_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralMatrixII(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]int{{1}}, generateMatrix(1))

	ast.Equal([][]int{
		{1, 2},
		{4, 3},
	}, generateMatrix(2))

	ast.Equal([][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}, generateMatrix(3))

	ast.Equal([][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}, generateMatrix(4))
}
