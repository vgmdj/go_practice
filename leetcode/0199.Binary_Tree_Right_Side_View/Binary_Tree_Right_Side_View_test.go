package Binary_Tree_Right_Side_View

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightSideView(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]int{1, 3, 4, 7}, rightSideView(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}))

}
