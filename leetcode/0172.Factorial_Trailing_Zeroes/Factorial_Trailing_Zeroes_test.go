package Factorial_Trailing_Zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorialTrailingZeroes(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, trailingZeroes(10))
	ast.Equal(3, trailingZeroes(16))

}
