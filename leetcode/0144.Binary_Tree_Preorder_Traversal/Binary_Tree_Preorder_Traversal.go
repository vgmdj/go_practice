package Binary_Tree_Preorder_Traversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	results := make([]int, 0)
	stack := list.New()
	p := root

	for p != nil || stack.Len() != 0 {
		for p != nil {
			results = append(results, p.Val)
			stack.PushBack(p)
			p = p.Left
		}

		if stack.Len() != 0 {
			node := stack.Back()
			stack.Remove(stack.Back())

			p = node.Value.(*TreeNode).Right
		}

	}

	return results

}
