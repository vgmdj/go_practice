package Additive_Number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAdditiveNumber(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(true, isAdditiveNumber("000"))
	ast.Equal(true, isAdditiveNumber("101"))
	ast.Equal(true, isAdditiveNumber("123"))
	ast.Equal(true, isAdditiveNumber("11235"))
	ast.Equal(true, isAdditiveNumber("199100199"))
	ast.Equal(false, isAdditiveNumber("112456712"))

}
