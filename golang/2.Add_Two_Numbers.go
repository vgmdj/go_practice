package main

import "log"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3, curr *ListNode
	var carry = 0

	for l1 != nil || l2 != nil {
		var value1, value2 int

		if l1 != nil {
			value1 = l1.Val
			l1 = l1.Next
		} else {
			value1 = 0
		}

		if l2 != nil {
			value2 = l2.Val
			l2 = l2.Next
		} else {
			value2 = 0
		}

		newNode := &ListNode{
			Val:  (value1 + value2 + carry) % 10,
			Next: nil,
		}
		carry = (value1 + value2 + carry) / 10

		if curr == nil {
			l3 = newNode
			curr = newNode
		} else {
			curr.Next = newNode
			curr = newNode
		}

		if carry > 0 {
			curr.Next = &ListNode{
				Val:  carry,
				Next: nil,
			}
		}

	}

	return l3
}

func main() {
	//l1 := ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//
	//l2 := ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 6,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}

	l11 := ListNode{
		Val:  1,
		Next: nil,
	}

	l22 := ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	//l3 := addTwoNumbers(&l1, &l2)
	//log.Println(l3.Val, l3.Next.Val, l3.Next.Next.Val)

	l33 := addTwoNumbers(&l11, &l22)
	log.Println(l33.Val, l33.Next.Val, l33.Next.Next.Val)

}
