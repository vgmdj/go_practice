package Longest_Increasing_Subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLIS(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1, lengthOfLIS([]int{2}))
	ast.Equal(1, lengthOfLIS([]int{2,2}))
	ast.Equal(6, lengthOfLIS([]int{1,3,6,7,9,4,10,5,6}))
	ast.Equal(5, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18,112}))

}
