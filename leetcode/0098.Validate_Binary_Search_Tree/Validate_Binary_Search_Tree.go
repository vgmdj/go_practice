package Validate_Binary_Search_Tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(q *TreeNode) bool {
	if q == nil || (q.Left == nil && q.Right == nil) {
		return true
	}
	return isValidTree(q, math.MaxInt64, math.MinInt64)
}

func isValidTree(q *TreeNode, max int, min int) bool {
	if q == nil {
		return true
	}
	if q.Val > min && q.Val < max {
		return isValidTree(q.Left, q.Val, min) && isValidTree(q.Right, max, q.Val)
	}
	return false
}
