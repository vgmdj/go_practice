package Reverse_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

//the first and better answer
func reverseList(head *ListNode) *ListNode {
	return reverseListRecursive(head)
}

func reverseListRecursive(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}
	tail := reverseListRecursive(node.Next)
	head := node
	h := head.Next
	h.Next = head
	head.Next = nil
	return tail
}

//the second answer
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}

//non-recursive answer
func reverseList3(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}

	p := head.Next
	for p.Next != nil{
		t := p.Next.Next
		p.Next.Next = head.Next
		head.Next = p.Next
		p.Next = t
	}

	node := head.Next
	p.Next = head
	head.Next = nil

	return node


}
