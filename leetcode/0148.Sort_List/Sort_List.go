package Sort_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	fast = slow.Next
	slow.Next = nil

	first, second := sortList(head), sortList(fast)
	return merge(first, second)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	head := &ListNode{}
	c := head

	for left != nil && right != nil {
		if left.Val < right.Val {
			c.Next = left
			left = left.Next
			c = c.Next

		} else {
			c.Next = right
			right = right.Next
			c = c.Next
		}
	}

	if left != nil {
		c.Next = left

	} else {
		c.Next = right
	}

	return head.Next
}

func sortList2(head *ListNode) *ListNode {
	r := &ListNode{}
	h := &ListNode{Next: head}

	for h.Next != nil {
		c := r

		for c.Next != nil {
			if c.Next.Val > h.Next.Val {
				break
			}

			c = c.Next
		}

		t := h.Next
		h.Next = t.Next
		t.Next = c.Next
		c.Next = t

	}

	return r.Next

}
