package Linked_List_Components

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool, len(G))
	for _, v := range G {
		m[v] = true
	}

	result := 1
	for head != nil {
		if !m[head.Val] {
			result++
		}

		head = head.Next

	}

	return result
}
