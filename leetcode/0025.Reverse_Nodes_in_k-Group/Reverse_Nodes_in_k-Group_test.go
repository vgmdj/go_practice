package Reverse_Nodes_in_k_Group

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{
				newList([]int{1, 2, 3, 4, 5, 6}),
				2,
			},
			newList([]int{2, 1, 4, 3, 6, 5}),
		},
		{
			"test2",
			args{
				newList([]int{1, 2, 3, 4, 5, 6, 7, 8}),
				3,
			},
			newList([]int{3, 2, 1, 6, 5, 4, 7, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newList(nums []int) *ListNode {
	Head := &ListNode{}
	tail := Head
	for _, v := range nums {
		node := &ListNode{
			Val: v,
		}

		tail.Next = node
		tail = tail.Next

	}

	return Head.Next

}
