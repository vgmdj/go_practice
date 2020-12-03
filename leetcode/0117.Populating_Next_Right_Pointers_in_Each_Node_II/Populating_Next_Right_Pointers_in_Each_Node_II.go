package Populating_Next_Right_Pointers_in_Each_Node_II

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(node *Node) {
			if node == nil {
				return
			}

			if nextStart == nil {
				nextStart = node
			}

			if last != nil {
				last.Next = node
			}
			last = node
		}

		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}

		start = nextStart

	}

	return root
}
