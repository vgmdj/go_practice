package Binary_Tree_Level_Order_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	levelOrderHelper(root, 0, &result)

	return result

}

func levelOrderHelper(root *TreeNode, level int, result *[][]int) {
	if root == nil{
		return
	}

	if len(*result) > level {
		(*result)[level] = append((*result)[level], root.Val)
	} else {
		*result = append(*result, []int{root.Val})
	}

	levelOrderHelper(root.Left, level+1, result)
	levelOrderHelper(root.Right, level+1, result)
}
