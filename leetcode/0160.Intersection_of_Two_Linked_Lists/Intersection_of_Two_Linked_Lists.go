package Intersection_of_Two_Linked_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB

		} else {
			p = p.Next

		}

		if q == nil {
			q = headA

		} else {
			q = q.Next

		}
	}

	return p
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var nodes = make(map[*ListNode]bool)
	for headA != nil {
		nodes[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if nodes[headB] {
			return headB
		}

		headB = headB.Next

	}

	return nil

}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lenA, lenB := 0, 0
	for p := headA; p != nil; p = p.Next {
		lenA++
	}

	for p := headB; p != nil; p = p.Next {
		lenB++
	}

	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	for lenA > lenB {
		headA = headA.Next
		lenA--
	}

	for headA != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next

	}

	return nil

}
