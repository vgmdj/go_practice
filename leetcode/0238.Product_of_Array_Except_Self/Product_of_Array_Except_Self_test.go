package Product_of_Array_Except_Self

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]int{2, 1}, productExceptSelf([]int{1, 2}))
	ast.Equal([]int{120, 60, 40, 30, 24}, productExceptSelf([]int{1, 2, 3, 4, 5}))

}
