package Unique_Binary_Search_Trees_II

import (
	"reflect"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
		{
			"case",
			args{
				2,
			},
			[]*TreeNode{
				&TreeNode{
					1,
					nil,
					&TreeNode{
						2,
						nil,
						nil,
					},
				},
				{
					2,
					&TreeNode{
						1,
						nil,
						nil,
					},
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
