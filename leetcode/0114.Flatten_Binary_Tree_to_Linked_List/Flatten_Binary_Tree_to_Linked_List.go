package Flatten_Binary_Tree_to_Linked_List

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	if root.Left == nil {
		return
	}

	leftTail := root.Left

	for leftTail.Right != nil {
		leftTail = leftTail.Right
	}

	leftTail.Right = root.Right
	root.Right = root.Left
	root.Left = nil

}
