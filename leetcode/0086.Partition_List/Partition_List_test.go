package Partition_List

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, partition(&ListNode{Val: 2, Next: &ListNode{Val: 1}}, 2))
	ast.Equal(&ListNode{Val: 1}, partition(&ListNode{Val: 1}, 0))
	ast.Equal(&ListNode{Val: 1}, partition(&ListNode{Val: 1}, 2))

	ast.Equal(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}, partition(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}, 3))

}
