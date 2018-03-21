package Add_Two_Numbers

import "testing"

func Test(t *testing.T) {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := addTwoNumbers(&l1, &l2)
	t.Log(l3.Val, l3.Next.Val, l3.Next.Next.Val)

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

	l33 := addTwoNumbers(&l11, &l22)
	t.Log(l33.Val, l33.Next.Val, l33.Next.Next.Val)
}
