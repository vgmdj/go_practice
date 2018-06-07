package Path_Sum

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return sumNode(root, 0, sum)

}

func sumNode(root *TreeNode, preSum int, sum int) bool {
	if root == nil {
		return false
	}

	preSum += root.Val

	if root.Left == nil && root.Right == nil {
		return preSum == sum
	}

	return sumNode(root.Left, preSum, sum) || sumNode(root.Right, preSum, sum)

}
