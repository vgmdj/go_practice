package _019_Remove_Nth_Node_From_End_of_List

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNode(t *testing.T) {
	ast := assert.New(t)

	list := &ListNode{Val: 1}
	AddNewNode(list, 2)
	AddNewNode(list, 3)
	AddNewNode(list, 4)

	except := &ListNode{Val: 1}
	AddNewNode(except, 2)
	AddNewNode(except, 4)
	ast.Equal(except, removeNthFromEnd(list, 2))

	AddNewNode(list, 5)
	AddNewNode(list, 6)
	AddNewNode(list, 7)

	AddNewNode(except, 6)
	AddNewNode(except, 7)
	ast.Equal(except, removeNthFromEnd(list, 3))

}

func AddNewNode(list *ListNode, v int) {
	p := list
	for p.Next != nil {
		p = p.Next
	}

	p.Next = &ListNode{
		Val: v,
	}

}
