package Balanced_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := treeDepth(root.Left)
	right := treeDepth(root.Right)

	if left > right+1 || left < right-1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := treeDepth(root.Left)
	right := treeDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}
