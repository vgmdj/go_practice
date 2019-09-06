package Reverse_Nodes_in_k_Group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	length := 0
	count := head
	for count != nil {
		count = count.Next
		length++
	}

	result := &ListNode{Next: head}
	c := result
	for i := k-1; i < length; i += k {
		c = reverseK(c, k)
	}

	return result.Next
}

func reverseK(result *ListNode, k int) *ListNode {
	c := result.Next

	for k > 1 && c.Next != nil {
		cn := c.Next
		c.Next = cn.Next
		cn.Next = result.Next
		result.Next = cn
		k--
	}

	return c

}
