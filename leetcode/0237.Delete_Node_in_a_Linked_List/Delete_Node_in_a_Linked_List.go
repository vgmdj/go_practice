package Delete_Node_in_a_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}

	node.Val = node.Next.Val
	if node.Next.Next == nil {
		node.Next = nil
		return
	}

	deleteNode(node.Next)

}
