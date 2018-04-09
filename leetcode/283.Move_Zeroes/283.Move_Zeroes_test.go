package Move_Zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	ast := assert.New(t)

	test1 := []int{0, 0, 1}
	moveZeroes(test1)
	checkArray(ast, test1, []int{1, 0, 0})

	test2 := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	moveZeroes(test2)
	checkArray(ast, test2, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0})

	test3 := []int{0, 0}
	moveZeroes(test3)
	checkArray(ast, test3, []int{0, 0})
}


func TestMoveZeroesII(t *testing.T) {
	ast := assert.New(t)

	test1 := []int{0, 0, 1}
	moveZeroesII(test1)
	checkArray(ast, test1, []int{1, 0, 0})

	test2 := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	moveZeroesII(test2)
	checkArray(ast, test2, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0})

	test3 := []int{0, 0}
	moveZeroesII(test3)
	checkArray(ast, test3, []int{0, 0})
}



func checkArray(ast *assert.Assertions, src, dest []int) {
	ast.Equal(len(src), len(dest))

	for k, v := range src {
		ast.Equal(v, dest[k])
	}

}
