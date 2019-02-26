package Partition_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var before, after = &ListNode{Next: head}, &ListNode{}
	var p, q = before, after
	for p.Next != nil {
		if p.Next.Val < x {
			p = p.Next
			continue
		}

		q.Next = p.Next
		p.Next = p.Next.Next
		q.Next.Next = nil
		q = q.Next

	}

	p.Next = after.Next

	return before.Next

}
