package Jump_Game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(false, canJump([]int{0, 2, 3}))
	ast.Equal(true, canJump([]int{2, 3, 1, 1, 4}))
	ast.Equal(false, canJump([]int{3, 2, 1, 0, 4}))
	ast.Equal(true, canJump([]int{6, 3, 2, 1, 0, 4}))

}
