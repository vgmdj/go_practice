package Consecutive_Numbers_Sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1,consecutiveNumbersSum(1))
	ast.Equal(1,consecutiveNumbersSum(2))
	ast.Equal(2,consecutiveNumbersSum(3))
	ast.Equal(2,consecutiveNumbersSum(5))
	ast.Equal(3,consecutiveNumbersSum(9))
	ast.Equal(4,consecutiveNumbersSum(15))
	ast.Equal(1,consecutiveNumbersSum(16))
	ast.Equal(4,consecutiveNumbersSum(30))
	ast.Equal(4, consecutiveNumbersSum(855877922))
}
