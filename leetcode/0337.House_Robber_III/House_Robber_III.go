package House_Robber_III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	v := calc(root)
	return max(v.include, v.exclude)
}

type value struct {
	include int
	exclude int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calc(root *TreeNode) value {
	if root == nil {
		return value{}
	}

	left := calc(root.Left)
	right := calc(root.Right)
	v := value{}
	v.include = root.Val + left.exclude + right.exclude
	v.exclude = max(left.include, left.exclude) + max(right.include, right.exclude)
	return v
}
