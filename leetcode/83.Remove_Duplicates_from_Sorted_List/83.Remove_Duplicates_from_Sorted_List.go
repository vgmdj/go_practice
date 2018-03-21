package Remove_Duplicates_from_Sorted_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for p := head; p.Next != nil; {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
			continue
		}

		p = p.Next
	}

	return head
}
