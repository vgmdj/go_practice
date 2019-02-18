package _024_Swap_Nodes_in_Pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapNodes(t *testing.T) {
	ast := assert.New(t)

	list := &ListNode{Val: 1}
	except := &ListNode{Val: 1}
	ast.Equal(except, swapPairs(list))

	list = &ListNode{Val:1}
	AddNewNode(list, 2)
	except = &ListNode{Val:2}
	AddNewNode(except, 1)
	ast.Equal(except, swapPairs(list))

	list = &ListNode{Val:2}
	AddNewNode(list,1)
	AddNewNode(list, 4)
	AddNewNode(list, 3)
	except = &ListNode{Val:1}
	AddNewNode(except,2)
	AddNewNode(except, 3)
	AddNewNode(except, 4)
	ast.Equal(except, swapPairs(list))

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
