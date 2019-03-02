package Set_Matrix_Zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetMatrixZeroes(t *testing.T) {
	ast := assert.New(t)

	case1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 0, 8},
		{0, 1, 2, 3},
	}
	setZeroes(case1)
	ast.Equal([][]int{
		{0, 2, 0, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}, case1)
}
