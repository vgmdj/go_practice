package SqrtX

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqrt(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(mySqrt(4), 2)
	ast.Equal(mySqrt(8), 2)
	ast.Equal(mySqrt(9), 3)
	ast.Equal(mySqrt(1), 1)
	ast.Equal(mySqrt(0), 0)

}
