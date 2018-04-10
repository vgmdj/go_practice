package Max_Consecutive_Ones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}), 3)
	ast.Equal(findMaxConsecutiveOnes([]int{0, 0, 0, 1, 1}), 2)
	ast.Equal(findMaxConsecutiveOnes([]int{0, 1}), 1)
	ast.Equal(findMaxConsecutiveOnes([]int{0}), 0)

}
