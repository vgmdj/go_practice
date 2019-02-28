package Remove_Duplicates_from_Sorted_List_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(&ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
		},
	}, deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}))

}
