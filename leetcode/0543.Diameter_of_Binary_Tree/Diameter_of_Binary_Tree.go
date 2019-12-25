package Diameter_of_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	deep(root, &max)
	return max
}

func deep(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := deep(root.Left, max)
	right := deep(root.Right, max)

	if *max > left+right+1 {
		*max = left + right + 1
	}

	if left > right {
		return left + 1

	}

	return right + 1

}
