package Reverse_Linked_List_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	var d = &ListNode{Next: head}
	var p, q = d, d
	for i := 1; i <= n; i++ {
		if i < m {
			p = p.Next
		}
		q = q.Next
	}

	for p.Next != q {
		n := p.Next
		p.Next = n.Next
		n.Next = q.Next
		q.Next = n
	}

	return d.Next
}
