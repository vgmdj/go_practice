package Bitwise_AND_of_Numbers_Range

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeBit(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(4, rangeBitwiseAnd(5, 7))
	ast.Equal(0, rangeBitwiseAnd(0, 1))
	ast.Equal(0, rangeBitwiseAnd(4000000, 2147483646))

}
