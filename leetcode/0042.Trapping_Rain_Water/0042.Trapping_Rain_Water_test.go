package Trapping_Rain_Water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	ast := assert.New(t)

	test1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ast.Equal(6, trap(test1))

	ast.Equal(6, trap2(test1))

	ast.Equal(6, trap3(test1))
}
