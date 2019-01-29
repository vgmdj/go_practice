package House_Robber_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobber(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1, rob([]int{1}))
	ast.Equal(2, rob([]int{1, 2}))
	ast.Equal(5, rob([]int{1, 1, 5}))
	ast.Equal(4, rob([]int{1, 2, 3, 1}))
	ast.Equal(11, rob([]int{2, 7, 9, 3, 1}))

}
