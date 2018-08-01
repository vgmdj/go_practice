package Convert_Sorted_Array_to_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(0, len(nums)-1, nums)
}

func dfs(left, right int, nums []int) *TreeNode {
	var n = right - left + 1
	if n <= 0 {
		return nil
	}

	var m = n / 2
	tp := &TreeNode{
		Val: nums[left+m],
	}

	tp.Left = dfs(left, left+m-1, nums)
	tp.Right = dfs(left+m+1, right, nums)

	return tp
}
