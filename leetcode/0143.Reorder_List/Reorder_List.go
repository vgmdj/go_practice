package Reorder_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	midNode := findMid(head)
	h2 := midNode.Next
	midNode.Next = nil
	h2 = reverse(h2)
	merge(head, h2)

}

func findMid(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	c := head
	for c != nil && c.Next != nil {
		next := c.Next
		c.Next = next.Next
		next.Next = dummy.Next
		dummy.Next = next
	}

	return dummy.Next

}

func merge(h1, h2 *ListNode) {
	c1, c2 := h1, h2
	for c1 != nil && c2 != nil {
		n2 := c2.Next
		c2.Next = c1.Next
		c1.Next = c2

		c1 = c2.Next
		c2 = n2
	}

}
