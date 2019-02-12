package Validate_Binary_Search_Tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	ast := assert.New(t)

	root := &TreeNode{
		Val:   10,
		Left:  getNewNode(5),
		Right: getNewNode(15),
	}
	root.Right.Left = getNewNode(6)
	root.Right.Right = getNewNode(20)
	ast.Equal(false, isValidBST(root))

	root1 := &TreeNode{
		Val:   3,
		Left:  getNewNode(1),
		Right: getNewNode(5),
	}

	root1.Left.Left = getNewNode(0)
	root1.Left.Right = getNewNode(2)
	root1.Right.Left = getNewNode(4)
	root1.Right.Right = getNewNode(6)
	root1.Right.Right.Right = getNewNode(3)
	ast.Equal(false, isValidBST(root1))

	root2 := &TreeNode{
		Val:   2,
		Left:  getNewNode(1),
		Right: getNewNode(3),
	}
	ast.Equal(true, isValidBST(root2))

}

func getNewNode(val int) *TreeNode {
	tree := new(TreeNode)
	tree.Val = val

	return tree
}
