package Binary_Tree_Postorder_Traversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	results := make([]int, 0)
	stack := list.New()
	p := root
	lastVisit := new(TreeNode)

	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}

		if stack.Len() != 0 {
			node := stack.Back()
			if node.Value.(*TreeNode).Right != nil && node.Value.(*TreeNode).Right != lastVisit {
				p = node.Value.(*TreeNode).Right
				continue
			}
			stack.Remove(node)

			results = append(results, node.Value.(*TreeNode).Val)
			lastVisit = node.Value.(*TreeNode)

		}

	}

	return results

}
