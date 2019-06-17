package Insertion_Sort_List

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

	p := &ListNode{Val: math.MinInt32, Next: head}

	pre, c := p, p.Next

	for c != nil {
		if c.Val >= pre.Val {
			c = c.Next
			pre = pre.Next
			continue
		}

		h := p
		hn := h.Next
		for hn != c {
			if c.Val <= hn.Val {
				pre.Next = c.Next
				c.Next = hn
				h.Next = c
				c = pre.Next
				break
			}

			h = h.Next
			hn = hn.Next
		}

	}

	return p.Next

}
