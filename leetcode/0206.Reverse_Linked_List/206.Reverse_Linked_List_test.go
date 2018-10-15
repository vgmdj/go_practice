package Reverse_Linked_List

import (
	"testing"
)

func TestReverseList(t *testing.T){

	list := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
				Next:&ListNode{
					Val:4,
					Next:&ListNode{
						Val:5,
						Next:nil,
					},
				},
			},
		},
	}

	t.Log(reverseList3(list))


}


