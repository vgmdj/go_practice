package Symmetric_Tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历并字符串
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var array []string
	getTreeArray(root, array)

	return true
}

func getTreeArray(root *TreeNode, array []string) {
	if root == nil {
		return
	}

	getTreeArray(root.Left, array)

	fmt.Println(root.Val)

	getTreeArray(root.Right, array)

}
