package Binary_Tree_Right_Side_View

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result []int
	helper(root, &result, 0)
	return result
}

func helper(root *TreeNode, result *[]int, level int) {
	if root == nil {
		return
	}

	if len(*result) == level {
		*result = append(*result, root.Val)
	}

	helper(root.Right, result, level+1)
	helper(root.Left, result, level+1)

}
