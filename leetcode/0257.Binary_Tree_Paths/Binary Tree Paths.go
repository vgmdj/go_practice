package Binary_Tree_Paths

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)

	dfsHelper(&result, []string{}, root)

	return result

}

func dfsHelper(results *[]string, subsets []string, root *TreeNode) {
	if root == nil {
		return
	}

	val := strconv.Itoa(root.Val)
	result := append([]string{}, subsets[:]...)
	result = append(result, val)

	if root.Left == nil && root.Right == nil {
		*results = append(*results, strings.Join(result, "->"))
		return

	}

	dfsHelper(results, result, root.Left)
	dfsHelper(results, result, root.Right)

}
