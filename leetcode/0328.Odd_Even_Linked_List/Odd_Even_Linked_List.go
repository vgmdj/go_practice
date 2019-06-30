package Odd_Even_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p, q := head, head.Next

	for q != nil && q.Next != nil {
		p.Next = q.Next
		q.Next = q.Next.Next
		p.Next.Next = q

		p = p.Next
		q = q.Next

	}

	return head
}

func oddEvenList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next

	}

	odd.Next = evenHead

	return head

}
