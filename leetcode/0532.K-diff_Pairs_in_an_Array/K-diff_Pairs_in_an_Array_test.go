package K_diff_Pairs_in_an_Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPairs(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1, findPairs([]int{1, 1, 1, 2, 2}, 1))
	ast.Equal(1, findPairs([]int{1, 1, 1, 1, 1, 1}, 0))
	ast.Equal(2, findPairs([]int{3, 1, 4, 1, 5}, 2))
	ast.Equal(2, findPairs([]int{1, 2, 1, 4, 5, 2}, 0))

}
