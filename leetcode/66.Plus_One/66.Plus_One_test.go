package Plus_One

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	ast.Equal(plusOne([]int{1, 2, 9}), []int{1, 3, 0})
	ast.Equal(plusOne([]int{9, 9, 9}), []int{1, 0, 0, 0})

}
