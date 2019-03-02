package Rotate_List

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateList(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(&ListNode{}, rotateRight(&ListNode{}, 4))
	ast.Equal(NewList([]int{4, 5, 1, 2, 3}), rotateRight(NewList([]int{1, 2, 3, 4, 5}), 2))
	ast.Equal(NewList([]int{2, 0, 1}), rotateRight(NewList([]int{0, 1, 2}), 4))
	ast.Equal(NewList([]int{1, 2}), rotateRight(NewList([]int{1, 2}), 2))
}

func NewList(values []int) *ListNode {
	list := &ListNode{}
	var p = list

	for k := range values {
		node := &ListNode{
			Val: values[k],
		}

		p.Next = node
		p = p.Next
	}

	return list.Next

}
