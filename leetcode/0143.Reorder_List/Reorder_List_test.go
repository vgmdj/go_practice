package Reorder_List

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reorderList(t *testing.T) {
	ast := assert.New(t)
	h1 := combineListFromArray([]int{1, 2, 3, 4, 5, 6})
	reorderList(h1)
	ast.Equal([]int{1, 6, 2, 5, 3, 4}, getArrayFromList(h1))

	h2 := combineListFromArray([]int{1, 2, 3, 4, 5})
	reorderList(h2)
	ast.Equal([]int{1, 5, 2, 4, 3}, getArrayFromList(h2))
}

func combineListFromArray(array []int) *ListNode {
	head := &ListNode{}
	c := head
	for _, val := range array {
		c.Next = &ListNode{
			Val: val,
		}
		c = c.Next
	}

	return head.Next

}

func getArrayFromList(head *ListNode) []int {
	result := make([]int, 0)
	c := head
	for c != nil {
		result = append(result, c.Val)
		c = c.Next
	}

	return result

}
