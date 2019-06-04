package _397_Integer_Replacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerReplacement(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0, integerReplacement(1))
	ast.Equal(1, integerReplacement(2))
	ast.Equal(2, integerReplacement(3))
	ast.Equal(2, integerReplacement(4))
	ast.Equal(3, integerReplacement(5))
	ast.Equal(3, integerReplacement(6))
	ast.Equal(4, integerReplacement(7))
	ast.Equal(3, integerReplacement(8))
	ast.Equal(4, integerReplacement(9))
	ast.Equal(4, integerReplacement(10))
	ast.Equal(5, integerReplacement(11))
	ast.Equal(4, integerReplacement(12))
	ast.Equal(5, integerReplacement(15))


}
