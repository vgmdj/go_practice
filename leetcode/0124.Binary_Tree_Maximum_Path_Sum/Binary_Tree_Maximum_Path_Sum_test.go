package Binary_Tree_Maximum_Path_Sum

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	if 6 != maxPathSum(tree1) {
		t.Error("error")
		return
	}

	tree2 := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	if 42 != maxPathSum(tree2) {
		t.Error("case2 error")
		return
	}

}
