package Populating_Next_Right_Pointers_in_Each_Node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.Next = root
	helper(root, nil)
	return root
}

func helper(left, right *Node) {
	if left == nil || left.Next == right {
		return
	}

	left.Next = right
	helper(left.Left, left.Right)
	if right == nil {
		return
	}
	helper(left.Right, right.Left)
	helper(right.Left, right.Right)

}

/*
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	// 每次循环从该层的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := leftmost; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right

			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	// 返回根节点
	return root
}
*/
