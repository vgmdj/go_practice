package Contains_Duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOK(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(containsDuplicate([]int{1, 2, 2, 3, 4, 5}), true)
	ast.Equal(containsDuplicate([]int{1, 8, 2, 3, 4, 5}), false)
	ast.Equal(containsDuplicate([]int{1, 1}), true)
	ast.Equal(containsDuplicate([]int{1, 1, 1, 1, 1}), true)

}
