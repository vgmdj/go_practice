package Range_Sum_Query_Immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumQuery(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	ast.Equal(1, obj.SumRange(0, 2))
	ast.Equal(1, obj.SumRange(0, 2))

	ast.Equal(-1, obj.SumRange(2, 5))

}
