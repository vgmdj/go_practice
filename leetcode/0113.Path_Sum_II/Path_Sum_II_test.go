package Path_Sum_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	ast := assert.New(t)

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	ast.Equal([][]int{{1, 2}}, pathSum(root, 3))
	ast.Equal([][]int{{1, 3}}, pathSum(root, 4))

}
