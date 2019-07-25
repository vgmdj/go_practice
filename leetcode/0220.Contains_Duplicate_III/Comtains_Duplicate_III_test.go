package Contains_Duplicate_III

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicates(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(false, containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9}, 2, 3))
	ast.Equal(true, containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	ast.Equal(false, containsNearbyAlmostDuplicate([]int{-1, 2147483647}, 1, 2147483647))

}
