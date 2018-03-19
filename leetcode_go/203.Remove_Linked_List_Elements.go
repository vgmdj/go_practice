package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	var node ListNode
	var p = &node

	for p.Next = head; p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return node.Next

}
