package Reverse_Linked_List_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(NewList([]int{1, 4, 3, 2, 5}), reverseBetween(NewList([]int{1, 2, 3, 4, 5}), 2, 4))
}

func NewList(l []int) *ListNode {
	var list ListNode
	var p = &list
	for _, v := range l {
		node := &ListNode{
			Val: v,
		}
		p.Next = node
		p = p.Next
	}

	return list.Next
}
