package Binary_Tree_Inorder_Traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInOrder(t *testing.T) {
	ast := assert.New(t)

	//[3,2,3,null,3,null,1]
	tree1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	ast.EqualValues([]int{2, 3, 3, 3, 1}, inorderTraversal(tree1))
}
