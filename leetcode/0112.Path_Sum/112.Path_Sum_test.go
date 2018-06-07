package Path_Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
the test case is :

			 3
		   /   \
         13	    5
        / \    / \
      nil 27  16 38


*/

func TestPathSum(t *testing.T) {
	ast := assert.New(t)

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  13,
			Left: nil,
			Right: &TreeNode{
				Val:   27,
				Left:  nil,
				Right: nil,
			},
		},

		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   16,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   38,
				Left:  nil,
				Right: nil,
			},
		},
	}

	ast.Equal(hasPathSum(root, 43), true)
	ast.Equal(hasPathSum(root, 24), true)
	ast.Equal(hasPathSum(root, 8), false)
	ast.Equal(hasPathSum(root, 27), false)

}
