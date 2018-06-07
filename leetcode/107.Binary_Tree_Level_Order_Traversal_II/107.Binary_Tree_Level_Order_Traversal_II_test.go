package Binary_Tree_Level_Order_Traversal_II

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

func TestTraversal(t *testing.T) {
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

	result := levelOrderBottom(root)

	ast.Equal(result[0][0], 27)
	ast.Equal(result[0][1], 16)
	ast.Equal(result[0][2], 38)
	ast.Equal(result[1][0], 13)
	ast.Equal(result[1][1], 5)
	ast.Equal(result[2][0], 3)

}
