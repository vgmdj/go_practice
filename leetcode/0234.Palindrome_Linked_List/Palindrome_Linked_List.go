package Palindrome_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var d = &ListNode{Next: head}
	slow, fast := d, d
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	last := reverse(slow.Next)
	for p := head; last != nil && p != nil; p = p.Next {
		if p.Val != last.Val {
			return false
		}

		last = last.Next
	}

	return true

}

func reverse(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}

	var d = &ListNode{Next: root}
	var p, q = d.Next, d.Next.Next

	for q != nil {
		p.Next = q.Next
		q.Next = d.Next
		d.Next = q
		q = p.Next
	}

	return d.Next
}
