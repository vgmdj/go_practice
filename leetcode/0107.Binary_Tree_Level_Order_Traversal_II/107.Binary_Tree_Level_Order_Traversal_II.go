package Binary_Tree_Level_Order_Traversal_II

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var rec = make(map[int][]int)
	treeRecord(root, 0, rec)

	var result [][]int
	for i := len(rec) - 1; i >= 0; i-- {
		result = append(result, rec[i])
	}

	return result
}

func treeRecord(root *TreeNode, level int, rec map[int][]int) {
	if root == nil {
		return
	}

	rec[level] = append(rec[level], root.Val)

	level++

	treeRecord(root.Left, level, rec)
	treeRecord(root.Right, level, rec)
}
