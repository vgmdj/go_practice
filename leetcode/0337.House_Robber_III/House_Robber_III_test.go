package House_Robber_III

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
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
	ast.Equal(7, rob(tree1))

}
