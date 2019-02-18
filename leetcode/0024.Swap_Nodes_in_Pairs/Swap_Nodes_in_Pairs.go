package _024_Swap_Nodes_in_Pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var list ListNode
	list.Next = head

	var p, q, position = &list, &list, 1

	for q != nil && q.Next != nil {
		if position%2 == 1 {
			q = q.Next
			position++
			continue
		}

		p.Next = q.Next
		q.Next = q.Next.Next
		p.Next.Next = q

		p = q
		position++

	}

	return list.Next
}
