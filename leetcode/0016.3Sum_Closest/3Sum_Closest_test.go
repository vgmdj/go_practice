package _3Sum_Closest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeNumClosest(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, threeSumClosest([]int{-1, 2, 1}, 1))
	ast.Equal(2, threeSumClosest([]int{-1, 2, 1}, 2))
	ast.Equal(2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	ast.Equal(-4, threeSumClosest([]int{-1, 2, 1, -4}, -4))
	ast.Equal(0, threeSumClosest([]int{-1, 2, -1, -4}, 0))
}
