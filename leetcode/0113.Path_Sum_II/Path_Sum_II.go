package Path_Sum_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	search(&result, []int{}, root, 0, sum)
	return result
}

func search(result *[][]int, subset []int, root *TreeNode, sum, target int) {
	if root == nil {
		return
	}

	if sum+root.Val == target && root.Left == nil && root.Right == nil {
		subset = append(subset, root.Val)
		*result = append(*result, append([]int{}, subset[:]...))
		return
	}

	search(result, append(subset, root.Val), root.Left, sum+root.Val, target)
	search(result, append(subset, root.Val), root.Right, sum+root.Val, target)

}
