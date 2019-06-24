package Convert_BST_to_Greater_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	helper(root, &sum)
	return root
}

func helper(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	helper(root.Right, sum)
	*sum += root.Val

	root.Val = *sum

	helper(root.Left, sum)

}
