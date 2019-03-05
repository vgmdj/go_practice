package Linked_List_Cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false

}

//use map
func hasCycle2(head *ListNode) bool {
	var p = head
	var visited = make(map[*ListNode]bool)

	for p != nil {
		if visited[p] {
			return true
		}

		visited[p] = true
		p = p.Next
	}

	return false

}
