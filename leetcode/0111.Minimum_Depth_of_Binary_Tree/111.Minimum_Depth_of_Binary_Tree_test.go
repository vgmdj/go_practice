package Minimum_Depth_of_Binary_Tree

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

func TestMiniDep(t *testing.T) {
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

	ast.Equal(minDepth(root), 3)

	root.Right.Right.Right = &TreeNode{
		Val:   54,
		Left:  nil,
		Right: nil,
	}

	ast.Equal(minDepth(root), 3)

	root.Left.Right = nil

	ast.Equal(minDepth(root), 2)

}
