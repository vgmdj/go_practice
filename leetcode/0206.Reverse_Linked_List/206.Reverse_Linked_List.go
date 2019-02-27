package Reverse_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

//the first and better answer
func reverseList(head *ListNode) *ListNode {
	return reverseListRecursive(head)
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}

//non-recursive answer
func reverseListNonRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var d = &ListNode{Next: head}
	p, c := head, head.Next
	for c != nil {
		p.Next = c.Next
		c.Next = d.Next
		d.Next = c
		c = p.Next
	}

	return d.Next

}
