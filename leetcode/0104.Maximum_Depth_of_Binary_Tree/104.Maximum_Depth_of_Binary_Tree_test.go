package Maximum_Depth_of_Binary_Tree

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

	log.Println(maxDepth(&root))
	log.Println(maxDepth(nil))
	log.Println(maxDepth(&TreeNode{0,nil,nil}))

}

func getNewNode(val int) *TreeNode {
	tree := new(TreeNode)
	tree.Val = val

	return tree
}
