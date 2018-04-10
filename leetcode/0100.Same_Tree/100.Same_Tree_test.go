package Same_Tree

import (
	"log"
	"testing"
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

	root2 := root

	log.Println(isSameTree(&root, &root2))

	log.Println(isSameTree(&root, nil))

}

func getNewNode(val int) *TreeNode {
	tree := new(TreeNode)
	tree.Val = val

	return tree
}
