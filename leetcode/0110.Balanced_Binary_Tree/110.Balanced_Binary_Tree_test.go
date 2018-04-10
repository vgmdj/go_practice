package Balanced_Binary_Tree

import (
	"testing"
	"log"
)

func TestOK(t *testing.T) {

	root := TreeNode{
		Val:   1,
		Left:  getNewNode(2),
		Right: getNewNode(2),
	}
	root.Left.Left = getNewNode(3)
	root.Left.Right = getNewNode(4)
	root.Right.Left = getNewNode(4)
	root.Right.Right = getNewNode(3)

	log.Println(isBalanced(&root))
	log.Println(isBalanced(nil))
	log.Println(isBalanced(&TreeNode{0,getNewNode(2),getNewNode(2)}))

}

func getNewNode(val int) *TreeNode {
	tree := new(TreeNode)
	tree.Val = val

	return tree
}
