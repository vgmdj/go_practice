package Unique_Paths_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))

	ast.Equal(0, uniquePathsWithObstacles([][]int{
		{0, 1, 0},
		{0, 0, 1},
		{0, 1, 0},
	}))

	ast.Equal(0, uniquePathsWithObstacles([][]int{
		{0, 0, 1, 0, 0},
	}))

}
