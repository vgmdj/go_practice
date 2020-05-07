package Split_Linked_List_in_Parts

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	p := root
	length := 0
	for p != nil {
		length++
		p = p.Next
	}
	mod := length % k
	size := length / k

	c := root
	result := make([]*ListNode, k)
	for i := 0; c != nil && i < k; i++ {
		result[i] = c
		curSize := 0
		if mod > 0 {
			curSize = size + 1
			mod--

		} else {
			curSize = size

		}

		for j := 1; j < curSize; j++ {
			c = c.Next

		}
		tail := c
		c = c.Next
		tail.Next = nil
	}
	return result
}
