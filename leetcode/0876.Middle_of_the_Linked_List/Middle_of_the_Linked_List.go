package Middle_of_the_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	tail, mid := head, head

	for tail != nil {
		mid = mid.Next
		tail = tail.Next
		if tail != nil {
			tail = tail.Next
		}
	}

	return mid

}
