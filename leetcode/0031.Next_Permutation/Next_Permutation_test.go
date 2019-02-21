package Next_Permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	ast := assert.New(t)

	case0 := []int{1, 1}
	nextPermutation(case0)
	ast.Equal([]int{1, 1}, case0)

	case1 := []int{1, 2, 3}
	nextPermutation(case1)
	ast.Equal([]int{1, 3, 2}, case1)

	case2 := []int{1}
	nextPermutation(case2)
	ast.Equal([]int{1}, case2)

	case3 := []int{}
	nextPermutation(case3)
	ast.Equal([]int{}, case3)

	case4 := []int{1, 1, 5}
	nextPermutation(case4)
	ast.Equal([]int{1, 5, 1}, case4)

	case5 := []int{1, 2, 5, 4, 3, 2, 1}
	nextPermutation(case5)
	ast.Equal([]int{1, 3, 1, 2, 2, 4, 5}, case5)

}
