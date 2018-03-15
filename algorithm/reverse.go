package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node

}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	log.Println(list.Val, list.Next.Val, list.Next.Next.Val)

	nl := reverse(list)

	log.Println(nl.Val, nl.Next.Val, nl.Next.Next.Val)

}
