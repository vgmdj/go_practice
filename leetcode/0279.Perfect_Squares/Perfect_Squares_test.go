package _279_Perfect_Squares

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerfectSquares(t *testing.T){
	ast := assert.New(t)


	ast.Equal(0,numSquares(0))
	ast.Equal(1,numSquares(1))
	ast.Equal(2,numSquares(2))
	ast.Equal(3,numSquares(3))
	ast.Equal(1,numSquares(4))
	ast.Equal(2,numSquares(5))
	ast.Equal(2,numSquares(8))
	ast.Equal(1,numSquares(9))
	ast.Equal(2,numSquares(10))
	ast.Equal(3,numSquares(12))
	ast.Equal(2,numSquares(13))
	ast.Equal(3,numSquares(14))
	ast.Equal(1,numSquares(16))
	ast.Equal(3,numSquares(19))
	ast.Equal(3,numSquares(48))
	ast.Equal(3,numSquares(235))



}