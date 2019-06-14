package Minimum_Size_Subarray_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSize(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
