package Binary_Tree_Inorder_Traversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result

}

func inorderTraversal2(root *TreeNode) []int {
	var result []int

	stack := list.New()
	p := root

	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}

		for p == nil && stack.Len() != 0 {
			node := stack.Back()
			stack.Remove(node)

			result = append(result, node.Value.(*TreeNode).Val)

			p = node.Value.(*TreeNode).Right

		}

	}

	return result

}
