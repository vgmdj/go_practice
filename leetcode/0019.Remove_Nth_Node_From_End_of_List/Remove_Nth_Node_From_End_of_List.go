package _019_Remove_Nth_Node_From_End_of_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var list ListNode
	list.Next = head

	var length int

	var p, q = &list, &list

	for q.Next != nil {
		if n == 0 {
			q = q.Next
			p = p.Next
		} else {
			q = q.Next
			n--
		}

		length++
	}

	if p.Next != nil {
		p.Next = p.Next.Next
		return list.Next
	}

	return nil

}
