package Gas_Station

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCan(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(3, canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	ast.Equal(-1, canCompleteCircuit([]int{3, 3, 4}, []int{3, 4, 4}))
	ast.Equal(4, canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))

}
