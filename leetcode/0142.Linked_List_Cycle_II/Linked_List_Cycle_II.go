package Linked_List_Cycle_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast != slow {
			continue
		}

		node := head
		for node != slow {
			node = node.Next
			slow = slow.Next
		}

		return node

	}

	return nil

}
