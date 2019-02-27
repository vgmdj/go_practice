package Reverse_Linked_List

import (
	"testing"

	"github.com/vgmdj/utils/logger"
)

func TestReverseList(t *testing.T) {

	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	show(reverseList(list))

}

func show(head *ListNode) {
	var p = head
	for p != nil {
		logger.Info(p.Val)
		p = p.Next
	}

}
