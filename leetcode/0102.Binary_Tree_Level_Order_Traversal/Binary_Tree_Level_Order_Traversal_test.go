package Binary_Tree_Level_Order_Traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	ast := assert.New(t)

	root := &TreeNode{
		Val:   3,
		Left:  getNewNode(9),
		Right: getNewNode(20),
	}

	root.Right.Left = getNewNode(15)
	root.Right.Right = getNewNode(7)

	ast.EqualValues([][]int{{3}, {9, 20}, {15, 7}}, levelOrder(root))

}

func getNewNode(val int) *TreeNode {
	node := new(TreeNode)
	node.Val = val

	return node
}
