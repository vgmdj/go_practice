package Contains_Duplicate_II

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOK(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(containsNearbyDuplicate([]int{1, 2, 3, 4, 5, 1}, 3), false)
	ast.Equal(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3), true)
	ast.Equal(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1), true)
	ast.Equal(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2), false)

}
