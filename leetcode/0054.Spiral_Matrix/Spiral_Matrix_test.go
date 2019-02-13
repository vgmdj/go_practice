package Spiral_Matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralMatrix(t *testing.T) {
	ast := assert.New(t)

	ast.EqualValues([]int{1, 2, 3, 4}, spiralOrder([][]int{
		{1}, {2}, {3}, {4}}))

	ast.EqualValues([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}))

	ast.EqualValues([]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}, spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}))

}
