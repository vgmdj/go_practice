package Maximal_Square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximal(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(4, maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))

}
