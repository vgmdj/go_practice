package _064_Minimum_Path_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(7, minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))

}
