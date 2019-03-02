package Maximum_Product_Subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0,maxProduct([]int{-2,0,-1}))
	ast.Equal(6, maxProduct([]int{1, 2, 3, 0, 4}))
}
