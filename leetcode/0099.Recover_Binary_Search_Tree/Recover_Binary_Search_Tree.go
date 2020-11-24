package Recover_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var x, y, pred, predecessor *TreeNode

	setPoint := func(c, p *TreeNode) {
		if p == nil || p.Val < c.Val {
			return
		}

		y = c
		if x == nil {
			x = p
		}
	}

	node := root
	for node != nil {
		if node.Left != nil {
			predecessor = node.Left
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				predecessor.Right = node
				node = node.Left
			} else {
				setPoint(node, pred)
				pred = node
				predecessor.Right = nil
				node = node.Right
			}

		} else {
			setPoint(node, pred)
			pred = node
			node = node.Right

		}

	}

	x.Val, y.Val = y.Val, x.Val

}
