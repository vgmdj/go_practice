package _090_Subsets_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetsII(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]int{}, subsetsWithDup([]int{1, 1, 2, 2}))

	ast.Equal([][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 2},
		{2},
		{2, 2},
	}, subsetsWithDup([]int{1, 2, 2}))

}
