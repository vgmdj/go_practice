package Recover_Binary_Search_Tree

import "testing"

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		},
		{
			"case",
			args{
				&TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
			if !sameTree(tt.args.root, tt.want) {
				t.Error("recover failed")
			}
		})
	}
}

func sameTree(r1, r2 *TreeNode) bool {
	if r1 == nil {
		if r2 == nil {
			return true
		}
		return false
	}

	if r2 == nil {
		return false
	}

	return r1.Val == r2.Val && sameTree(r1.Left, r2.Left) && sameTree(r1.Right, r2.Right)

}
