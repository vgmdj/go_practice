package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 ListNode
	var tmp = &l3

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tmp.Next = l2
			l2 = l2.Next
		} else {
			tmp.Next = l1
			l1 = l1.Next
		}

		tmp = tmp.Next
	}

	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}

	return l3.Next

}
