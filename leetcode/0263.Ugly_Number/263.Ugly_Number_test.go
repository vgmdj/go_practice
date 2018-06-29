package Ugly_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOK(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(isUgly(1), true)
	ast.Equal(isUgly(6), true)
	ast.Equal(isUgly(15), true)
	ast.Equal(isUgly(14), false)
	ast.Equal(isUgly(8), true)
	ast.Equal(isUgly(13), false)
	ast.Equal(isUgly(0), false)
	ast.Equal(isUgly(-2147483648), false)

}
