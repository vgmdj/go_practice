package Convert_Sorted_List_to_Binary_Search_Tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	ast := assert.New(t)

	list1 := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	}
	case1 := []int{-10, -3, 0, 5, 9}
	ast.EqualValues(case1, convertListToArray(list1))

	tree1 := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -10,
			Right: &TreeNode{
				Val: -3,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	ast.Equal(tree1, sortedArrayToBST(0, len(case1)-1, case1))

	ast.Equal(tree1, sortedListToBST(list1))

}
