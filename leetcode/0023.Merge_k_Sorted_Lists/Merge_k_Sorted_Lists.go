package Merge_k_Sorted_Lists

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil

	case 1:
		return lists[0]

	case 2:
		return mergeTwoLists(lists[0], lists[1])

	default:
		return mergeTwoLists(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))

	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := new(ListNode)
	tempNode := headNode

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tempNode.Next = l1
			l1 = l1.Next

		} else {
			tempNode.Next = l2
			l2 = l2.Next

		}

		tempNode = tempNode.Next
	}

	if l1 != nil {
		tempNode.Next = l1

	} else if l2 != nil {
		tempNode.Next = l2

	}

	return headNode.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	result := &ListNode{}
	c := result
	length := len(lists)

	for length != 0 {

		mink, minv := -1, math.MaxInt32
		for i := 0; i < length; i++ {
			if lists[i] == nil {
				mink = i
				break
			}

			if lists[i].Val < minv {
				mink = i
				minv = lists[i].Val
			}
		}

		if lists[mink] == nil {
			if mink != length-1 {
				lists[mink], lists[length-1] = lists[length-1], lists[mink]
			}

			length--
			continue
		}

		c.Next = lists[mink]
		lists[mink] = lists[mink].Next
		c = c.Next

	}

	return result.Next
}
