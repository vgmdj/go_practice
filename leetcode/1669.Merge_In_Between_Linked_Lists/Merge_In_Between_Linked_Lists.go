package Merge_In_Between_Linked_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	bp1, bp2 := list1, list1
	for c := 1; bp1 != nil && c < a; bp1, c = bp1.Next, c+1 {

	}

	bp2 = bp1.Next
	for c := a; bp2 != nil && c <= b; bp2, c = bp2.Next, c+1 {

	}

	head, tail := list2, list2
	for tail.Next != nil {
		tail = tail.Next
	}

	bp1.Next = head
	tail.Next = bp2

	return list1

}
