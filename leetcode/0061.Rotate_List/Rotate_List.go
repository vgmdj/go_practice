package Rotate_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	var p, q = head, head

	for i := 0; i < k; i++ {
		if q.Next == nil {
			return rotateRight(head, k%(i+1))
		}

		q = q.Next
	}

	for q.Next != nil {
		p = p.Next
		q = q.Next
	}

	newNode := p.Next
	q.Next = head
	p.Next = nil

	return newNode

}
