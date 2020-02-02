package Binary_Tree_Maximum_Path_Sum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	helper(root, &result)
	return result

}

func helper(root *TreeNode, result *int) int {
	if root == nil {
		return -1
	}

	left := helper(root.Left, result)
	right := helper(root.Right, result)
	mid := root.Val

	cMax := Max(mid, mid+left, mid+right)
	*result = Max(*result, cMax, mid+left+right)

	return cMax

}

func Max(nums ...int) int {
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		if r < nums[i] {
			r = nums[i]
		}
	}

	return r
}
