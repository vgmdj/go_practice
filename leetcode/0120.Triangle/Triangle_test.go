package _120_Triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangle(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(11, minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))

}
