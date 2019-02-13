package Sum_Root_to_Leaf_Numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}

	val = val*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return val
	}

	return helper(root.Left, val) + helper(root.Right, val)

}
