package Wiggle_Subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWiggle(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0,wiggleMaxLength([]int{}))
	ast.Equal(1, wiggleMaxLength([]int{0, 0}))
	ast.Equal(6, wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	ast.Equal(7, wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	ast.Equal(2, wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))

}
