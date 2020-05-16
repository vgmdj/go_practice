package Flatten_a_Multilevel_Doubly_Linked_List

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	head := new(Node)
	head.Next = root
	cur := head

	stack := make([]*Node, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		cur.Next = p
		p.Prev = cur

		if p.Next != nil {
			p.Next.Prev = nil
			stack = append(stack, p.Next)
			p.Next = nil
		}

		if p.Child != nil {
			p.Child.Prev = nil
			stack = append(stack, p.Child)
			p.Child = nil
		}

		cur = p

	}

	root.Prev = nil

	return root

}
