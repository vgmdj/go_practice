package Invert_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var left = invertTree(root.Left)
	var right = invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root

}
