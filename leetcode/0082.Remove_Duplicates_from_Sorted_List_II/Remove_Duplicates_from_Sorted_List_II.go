package Remove_Duplicates_from_Sorted_List_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var d = &ListNode{Next: head}
	var p, q = d, d.Next

	for q.Next != nil {
		if p.Next.Val == q.Next.Val {
			q = q.Next
			continue
		}

		if p.Next == q {
			p = p.Next

		} else {
			p.Next = q.Next
		}

		q = q.Next
	}

	if p.Next != q {
		p.Next = q.Next
	}

	return d.Next

}
