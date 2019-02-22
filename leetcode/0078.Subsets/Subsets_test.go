package Subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]int{
		[]int{},
		[]int{1},
	}, subsets([]int{1}))

	ast.Equal([][]int{
		[]int{},
		[]int{1},
		[]int{1, 3},
		[]int{1, 3, 5},
		[]int{1, 5},
		[]int{3},
		[]int{3, 5},
		[]int{5}},
		subsets([]int{1, 3, 5}))

}
