package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for p := head; p.Next != nil; {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
			continue
		}

		p = p.Next
	}

	return head
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	l = deleteDuplicates(l)

	log.Println(l.Val, l.Next.Val)

}
