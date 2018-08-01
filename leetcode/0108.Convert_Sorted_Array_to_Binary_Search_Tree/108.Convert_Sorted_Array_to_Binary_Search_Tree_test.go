package Convert_Sorted_Array_to_Binary_Search_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	tree
		0
      /  \
     -3  9
    /   /
  -10   5
*/

func TestOK(t *testing.T) {
	ast := assert.New(t)

	case1 := []int{-10, -3, 0, 5, 9}
	tree1 := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -10,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	ast.Equal(sortedArrayToBST(case1), tree1)

}
