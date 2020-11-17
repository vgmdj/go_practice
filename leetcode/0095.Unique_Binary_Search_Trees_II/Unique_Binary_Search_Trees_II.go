package Unique_Binary_Search_Trees_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{nil}

	combine := func(rootVal int, left, right []*TreeNode) []*TreeNode {
		result := make([]*TreeNode, 0)
		for _, lNode := range left {
			for _, rNode := range right {
				result = append(result, &TreeNode{
					rootVal,
					lNode,
					rNode,
				})
			}
		}
		return result
	}

	leftJoin := func(n int) []*TreeNode {
		return dp[n]
	}

	rightJoin := func(n int, offset int) []*TreeNode {
		result := make([]*TreeNode, 0)
		for _, node := range dp[n] {
			result = append(result, cloneTree(node, offset))
		}

		return result
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			left := leftJoin(j - 1)
			right := rightJoin(i-j, j)
			dp[i] = append(dp[i], combine(j, left, right)...)
		}
	}

	return dp[n]
}

func cloneTree(root *TreeNode, offset int) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{
		Val:   root.Val + offset,
		Left:  cloneTree(root.Left, offset),
		Right: cloneTree(root.Right, offset),
	}

}
